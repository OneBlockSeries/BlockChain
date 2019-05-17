package main

import(	
	
	
)


func main(){
	
	bl:=StartBlockChain()
	
	bl.Runing()
	
	bl.SendCoin("aa","bb",2)
	bl.SendCoin("cc","dd",4)
}