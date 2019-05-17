package main
import (
	"fmt"
)



type Blockchain struct{

	blocks []*Block
}

func StartBlockChain() (bl *Blockchain){


	b:=GenesisBlock()
	fmt.Printf("startblock=%x\n",b.bhead.hash)
	return &Blockchain{[]*Block{b}}
}

func (bl *Blockchain) Runing(){
	fmt.Printf("blockchain running\n")
	/*------加入到区块链-------*/
	tr:=Transaction{}
	bl.AddBlockToChain([]Transaction{tr})

	/*-------把blocks里的区块打印出来-----------*/
	for idx,value:=range bl.blocks{
		fmt.Printf("%d,block,hash=%x\n",idx,value.bhead.hash)
	}
}
func GenesisBlock() *Block{
	//message:="Genesis block"
	GenesTransaction:=Transaction{}
	Genesisblock:=CreatBlock([]Transaction{GenesTransaction},[]byte{})  //没有prevhash 第一个块，当然没有前向块的哈希
	return Genesisblock
}
func(bl *Blockchain)AddBlockToChain(trans  []Transaction){
	
	block:=CreatBlock(trans,bl.blocks[len(bl.blocks)-1].bhead.hash)
	bl.blocks=append(bl.blocks,block)
}

func (bl* Blockchain)SendCoin(from, to string,mount int){

	//封装成transaction

	fmt.Printf("sendcoin from=%s,to=%s,mount=%d\n",from,to,mount);
}