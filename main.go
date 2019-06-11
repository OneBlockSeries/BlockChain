package main

import ()

func main() {

	w1:=CreateWallet()
	w2:=CreateWallet()
	w3:=CreateWallet()
	w4:=CreateWallet()

	bl := StartBlockChain()
	bl.Runing()

	

	bl.SendCoin(string(w1.GetBase58Address()), string(w2.GetBase58Address()), 2)
	bl.SendCoin(string(w4.GetBase58Address()), string(w3.GetBase58Address()), 4)
}
