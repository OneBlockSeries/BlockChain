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
	bl.SendCoin(GenesisWallet.GetBase58Address(), w1.GetBase58Address(), 10)
	bl.SendCoin(GenesisWallet.GetBase58Address(), w2.GetBase58Address(), 10)
	bl.SendCoin(GenesisWallet.GetBase58Address(), w3.GetBase58Address(), 10)
	bl.SendCoin(GenesisWallet.GetBase58Address(), w3.GetBase58Address(), 10)

	//相互转化
	bl.SendCoin(w1.GetBase58Address(), w2.GetBase58Address(), 2)
	bl.SendCoin(w4.GetBase58Address(), w3.GetBase58Address(), 4)
	bl.SendCoin(w2.GetBase58Address(), w4.GetBase58Address(), 4)
}
