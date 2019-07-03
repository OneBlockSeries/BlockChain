package main

import (
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

func StartBlockChain(genesisAddress []byte) (bl *Blockchain) {

	b := GenesisBlock(genesisAddress)
	fmt.Printf("startblock=%x\n", b.bhead.hash)
	return &Blockchain{[]*Block{b}}
}

func (bl *Blockchain) Runing() {
	fmt.Printf("blockchain running\n")
	/*------加入到区块链-------*/

	tr := Transaction{}
	bl.AddBlockToChain([]*Transaction{&tr})

	/*-------把blocks里的区块打印出来-----------*/
	for idx, value := range bl.blocks {
		fmt.Printf("%d,block,hash=%x\n", idx, value.bhead.hash)
	}
}
func GenesisBlock(genesisAddress []byte) *Block {
	//message:="Genesis block"

	GenesTransaction := Coinbase(genesisAddress)
	Genesisblock := CreatBlock([]*Transaction{GenesTransaction}, []byte{}) //没有prevhash 第一个块，当然没有前向块的哈希
	
	return Genesisblock
}
func (bl *Blockchain) AddBlockToChain(trans []*Transaction) {

	block := CreatBlock(trans, bl.blocks[len(bl.blocks)-1].bhead.hash)
	bl.blocks = append(bl.blocks, block)
}

func (bl *Blockchain) SendCoin(from, to string, mount int) bool{

	//封装成transaction
	var ins []Input
	var outs []Output

	//判断address是否合法
	if IsTrueAddress(from)!=true||IsTrueAddress(to)!=true{
		fmt.Printf("sendcoin  wrong address\n")
		return false
	}

	//两步走 step1,从区块链中找from 找出所有的未花费的输出， 
	//step2，封装输入以找出来的输出填 当然得证明是自己的钱
	//封装输出  不要忘记有找零情况

	enoughOutput,findmount:=bl.FindEnoughOutputs(from)
	if findmount<mount{
		fmt.Printf("SEND NOT ENOUGH money,%d",enoughOutput)
		return false
	}

	//1.证明是自己的钱，解锁输出，封装输入，2.封装输出，考虑找零情况


	tr:=Transaction{nil,ins,outs}
	tr.SetId()
	fmt.Printf("sendcoin from=%s,to=%s,mount=%d\n", from, to, mount)
	return true
}
func (bl *Blockchain)FindEnoughOutputs(from string)([]Output,int){

	//从区块链中迭代找出住够多的钱
	return []Output{},1;
}