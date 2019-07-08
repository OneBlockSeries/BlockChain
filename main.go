package main

import (
	"fmt"
)

func main() {

	GenesisWallet:=CreateWallet()
	fmt.Printf("Geniswallet Addredd=%s\n",GenesisWallet.GetBase58Address())
	w1:=CreateWallet()
	fmt.Printf("w1----------Addredd=%s\n",w1.GetBase58Address())
	w2:=CreateWallet()
	fmt.Printf("w2----------Addredd=%s\n",w2.GetBase58Address())
	w3:=CreateWallet()
	fmt.Printf("w3----------Addredd=%s\n",w3.GetBase58Address())
	w4:=CreateWallet()
	fmt.Printf("w4----------Addredd=%s\n",w4.GetBase58Address())


	bl := StartBlockChain(GenesisWallet.GetBase58Address())
	//bl.Runing()
	fmt.Printf("start to send coin\n")
	//创世快 发送给w1,w2,w3,w4 各10个
	Gentr1:=bl.SendCoin(GenesisWallet.GetBase58Address(), w1.GetBase58Address(), 10,GenesisWallet)
	if Gentr1!=nil{
		bl.Runing(Gentr1)
	}
	Gentr2:=bl.SendCoin(GenesisWallet.GetBase58Address(), w2.GetBase58Address(), 10,GenesisWallet)
	if Gentr2!=nil{
		bl.Runing(Gentr2)
	}
	Gentr3:=bl.SendCoin(GenesisWallet.GetBase58Address(), w3.GetBase58Address(), 10,GenesisWallet)
	if Gentr3!=nil{
		bl.Runing(Gentr3)
	}
	Gentr4:=bl.SendCoin(GenesisWallet.GetBase58Address(), w3.GetBase58Address(), 10,GenesisWallet)
	if Gentr4!=nil{
		bl.Runing(Gentr4)
	}

	//相互转化
	tr12:=bl.SendCoin(w1.GetBase58Address(), w2.GetBase58Address(), 11,w1)
	if tr12!=nil{
		bl.Runing(tr12)
	}
	tr43:=bl.SendCoin(w4.GetBase58Address(), w3.GetBase58Address(), 4,w4)
	if tr43!=nil{
		bl.Runing(tr43)
	}
	tr24:=bl.SendCoin(w2.GetBase58Address(), w4.GetBase58Address(), 10,w2)
	if tr24!=nil{
		bl.Runing(tr24)
	}
}
