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
//输入  引用之前的输出，除了coinbase
type Input struct {
}

//输出  币存储地方
type Output struct {
	Value int
	Who   string //比特币里是有解锁脚本
}

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

/* coinbase交易，挖矿奖励，创世快奖励，没有输入input,只有输出output   */

func Coinbase(to string) *Transaction {
	fmt.Printf("coninbase\n")
	input := Input{}
	output := Output{rewardcoin, to}
	tran := Transaction{nil, []Input{input}, []Output{output}} //nil 不需要引用之前的输出，
	tran.SetId()
	return &tran
}
