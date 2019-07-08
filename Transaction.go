package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	//"encoding/hex"
	"fmt"
	"log"
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
	output := Output{rewardcoin, to}
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
/* //数据签名
func (tx *Transaction) Sign(privkey ecdsa.PrivateKey,prevTrans map[string]Transaction){
	if tx.isCoinbase(){
		return
	}

}
*/
func (tx *Transaction)isCoinbase()bool{

	return false
}