package main
import (
	"fmt"
)



type Blockchain struct{

	blocks []*Block
}

func StartBlockChain(message string) (bl *Blockchain){

	fmt.Printf("message=%s\n",message)
	b:=&Block{message}
	return &Blockchain{[]*Block{b}}
}

func (bl *Blockchain) Runing(){
	fmt.Printf("blockchain running\n")
}