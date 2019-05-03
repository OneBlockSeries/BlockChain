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
	for idx,value:=range bl.blocks{
		fmt.Printf("%d,block,hash=%x\n",idx,value.bhead.hash)
	}
}
func GenesisBlock() *Block{
	message:="Genesis block"
	Genesisblock:=CreatBlock(message,[]byte{})  //没有prevhash 第一个块，当然没有前向块的哈希
	return Genesisblock
}