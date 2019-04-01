package main

import(
	"fmt"
)

func main(){
	fmt.printf("start\n")
	block:=Block{}
	BLchain:=BlockChain{}
}

//----------Block----////

type BlockHead struct{			// 区块头
	timestamp  int32
	PrevBlockHash  []byte
	NowBlockHash   []byte
}
type Transaction struct{		//交易结构体

	message string
}
type Block struct{				//区块  区块=区块头+许多交易

	blockhead BlockHead
	TransactionS []Transaction
}


//-------BlockChain---///
type BlockChain struct{			//区块链=许多区块链接在一起

	Blocks  []Block
}