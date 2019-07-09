package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"crypto/ecdsa"
	//"encoding/hex"
	"fmt"
	"log"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

const rewardcoin = 50 //奖励钱

//输入  引用之前的输出， 
type Input struct {
	Id  []byte		//对应之前所在Transaction Id
	OutId int 		//具体哪个输出，所对应的Transaction有多个输出，比如找零,具体哪个
	//unlock string 	//证明你可以用这笔钱，也就是证明这是你的钱 ，数字签名理论
	Pubkey  []byte		//没有哈希的pubkey ,
	Signature []byte	//整个交易的一个签名
}
//输出  币存储地方
type Output struct {
	Value int
	Pubkeyhash   []byte //哈希过的  RIPEMD16(SHA256(PubKey))
}
func (in *Input)UsesKey(pubkeyhash []byte) bool{
	lockingHash:=PubkeyTwicehash(in.Pubkey)
	return bytes.Compare(lockingHash,pubkeyhash)==0
}
func (out *Output)Lock(address []byte){
	base58decode:=Base58Decode(address)
	pubkeyhash:=base58decode[1:len(base58decode)-4]
	out.Pubkeyhash=pubkeyhash
	//return out.Pubkeyhash==unlockingData
}
func (in *Input)CanUnlockOutput(address []byte)bool{
	
	return true;
}
func (out *Output) IsLocked(pubkeyhash []byte) bool{
	return bytes.Compare(out.Pubkeyhash,pubkeyhash)==0
}
func (out *Output) CanbeUnlockWith(address []byte)bool{
	return true
}
func (out *Output)CanBeUnlock(address []byte)bool{
	return true
}
///func AddressToPubkeyhash(address []byte)[]byte{

//}
/*--一笔交易就是包括一个多个输入和一个或多个输出打包成成一个Transaction,然后被挖矿挖出加入到区块链中---*/
type Transaction struct {
	Id   []byte
	Ins  []Input
	Outs []Output
}
func (tr *Transaction) CopyPartTransaction() *Transaction{
	var inputs []Input
	var outputs []Output
	for _,vin:=range tr.Ins{
		inputs=append(inputs,Input{vin.Id,vin.OutId,nil,nil}) //pubkey 和signage设置为nil
	}
	for _,vout:=range tr.Outs{
		outputs=append(outputs,Output{vout.Value,vout.Pubkeyhash})
	}
	txcopy:=&Transaction{tr.Id,inputs,outputs}
	return txcopy
}
func (tr *Transaction) SetId() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err :=enc.Encode(tr)
	
	if err != nil {
		log.Panic(err)
	}
	
	//把自身的数据做一次哈希
	hash = sha256.Sum256(encoded.Bytes())
	tr.Id = hash[:]
	fmt.Printf("the transaction hash=%x\n",tr.Id)
}

/* coinbase交易，特殊交易，挖矿奖励，创世快奖励，没有输入input,只有输出output   */

func Coinbase(to []byte) *Transaction {
	fmt.Printf("coninbase\n")
	input := Input{}
	output := Output{rewardcoin, []byte{}}
	output.Lock(to)
	tran := Transaction{nil, []Input{input}, []Output{output}} //nil 不需要引用之前的输出，
	tran.SetId()
	return &tran
}
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}
 //数据签名
func (tx *Transaction) Sign(privkey ecdsa.PrivateKey,prevTxs map[string]*Transaction){
	//1.用于签名数据 引用的输出里的公钥哈希(发送方)+产生的新输出里的公钥哈希(接收方)+新的输出值
	//2. 签名
	if tx.isCoinbase(){
		return
	}
	txCopy:=tx.CopyPartTransaction()
	
	for ID,vin:=range txCopy.Ins{

		prevTx:=prevTxs[BytesToString(vin.Id)]  //之前引用的交易
		txCopy.Ins[ID].Signature=nil   //重复，再次确认操作
		txCopy.Ins[ID].Pubkey=prevTx.Outs[vin.OutId].Pubkeyhash //把引用输出里的pubkeyhash 调出来
		txCopy.SetId()
		txCopy.Ins[ID].Pubkey=nil   //重置，不影响下一个

		r,s,_:=ecdsa.Sign(rand.Reader,&privkey,txCopy.Id)
		sig:=append(r.Bytes(),s.Bytes()...)
		tx.Ins[ID].Signature=sig

		txCopy.Id=[]byte{} //重置下
	}

}
//签名验证
func (tx *Transaction)Verify(prevTxs map[string]*Transaction)bool{
	//1.检查交易输入有权使用之前交易输出，存储的pubkey 哈希后与所引用输出哈希相匹配，保证了发送方只能花费自己的币
	//2 签名是正确的，保证了交易由币的实际拥有者创建
	txCopy := tx.CopyPartTransaction()
    curve := elliptic.P256()

    for inID, vin := range tx.Ins {
        prevTx := prevTxs[BytesToString(vin.Id)]
        txCopy.Ins[inID].Signature = nil
        txCopy.Ins[inID].Pubkey = prevTx.Outs[vin.OutId].Pubkeyhash
        txCopy.SetId()
        txCopy.Ins[inID].Pubkey = nil

        r := big.Int{}
        s := big.Int{}
        sigLen := len(vin.Signature)
        r.SetBytes(vin.Signature[:(sigLen / 2)])
        s.SetBytes(vin.Signature[(sigLen / 2):])

        x := big.Int{}
        y := big.Int{}
        keyLen := len(vin.Pubkey)
        x.SetBytes(vin.Pubkey[:(keyLen / 2)])
        y.SetBytes(vin.Pubkey[(keyLen / 2):])

        rawPubKey := ecdsa.PublicKey{curve, &x, &y}
        if ecdsa.Verify(&rawPubKey, txCopy.Id, &r, &s) == false {
            return false
		}
		txCopy.Id=[]byte{} //重置下
    }
	return true;
}


func (tx *Transaction)isCoinbase()bool{

	return false
}