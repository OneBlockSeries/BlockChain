package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"encoding/binary"
	"log"
)
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

var (
	maxNonce = int64(math.MaxInt64)
)

const targetBits = 16 //可以调整 16就相当于前两位是0   调整24就是前三位
type ProofOfWork struct {
	block  *Block
	target *big.Int //大整数，用于和区块哈希结果相比
}

func CreateProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)                  //定义1
	target.Lsh(target, uint(256-targetBits)) //右移动targetBits（16）位，000000....100000....前面16个是0，总共256位
	pow := &ProofOfWork{block, target}

	return pow
}

//准备数据
func (pow *ProofOfWork) PowprepareDate(nonce int64) (date []byte) {

	data := bytes.Join([][]byte{
		IntToHex(pow.block.bhead.timestamp),
		pow.block.bhead.prevhash,
		pow.block.bhead.hash,
		IntToHex(pow.block.bhead.nonce),
		pow.block.bhead.merkleRoot,
	}, []byte{})
	return data
}

func (pow *ProofOfWork) PowRun() (returnnonce int64, returnhash []byte) {

	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)
	for nonce < maxNonce {
		data := pow.PowprepareDate(nonce)
		hash  =sha256.Sum256(data)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("nonce is %d\n", nonce)

	return nonce, hash[:]

}

//验证
func (pow *ProofOfWork) IsValidate() bool {

	var hashInt big.Int
	data:=pow.PowprepareDate(pow.block.bhead.nonce)
	hash:=sha256.Sum256(data)
	hashInt.SetBytes(hash[:])


	return hashInt.Cmp(pow.target)==-1
}
