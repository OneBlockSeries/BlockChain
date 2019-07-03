package main

import (

)

func main() {

	GenesisWallet:=CreateWallet()
	w1:=CreateWallet()
	w2:=CreateWallet()
	w3:=CreateWallet()
	w4:=CreateWallet()


	bl := StartBlockChain(GenesisWallet.GetBase58Address())
	bl.Runing()

	//创世快 发送给w1,w2,w3,w4 各10个
	bl.SendCoin(string(GenesisWallet.GetBase58Address()), string(w1.GetBase58Address()), 10)
	bl.SendCoin(string(GenesisWallet.GetBase58Address()), string(w2.GetBase58Address()), 10)
	bl.SendCoin(string(GenesisWallet.GetBase58Address()), string(w3.GetBase58Address()), 10)
	bl.SendCoin(string(GenesisWallet.GetBase58Address()), string(w3.GetBase58Address()), 10)

	//相互转化
	bl.SendCoin(string(w1.GetBase58Address()), string(w2.GetBase58Address()), 2)
	bl.SendCoin(string(w4.GetBase58Address()), string(w3.GetBase58Address()), 4)
	bl.SendCoin(string(w2.GetBase58Address()), string(w4.GetBase58Address()), 4)
}
