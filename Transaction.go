package main

import(
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	//"encoding/hex"
	"log"
)

//输入  引用之前的输出，除了coinbase
type Input struct{

}
//输出  币存储地方
type Output struct{
 
	value  int 

}
/*--一笔交易就是包括一个多个输入和一个或多个输出打包成成一个Transaction,然后被挖矿挖出加入到区块链中---*/
type Transaction struct{
	Id    []byte
	ins   []Input
	outs  []Output
}
func (tr *Transaction) SetId(){
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tr)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tr.Id = hash[:]
}