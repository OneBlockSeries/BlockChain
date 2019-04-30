package main

import(	
	
	"fmt"
)


func main(){
	fmt.Printf("the first day \n")
	
	bl:=StartBlockChain("mybank")
	bl.Runing()
}