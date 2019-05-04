package main

import(
	
)

//输入
type Input struct{

}
//输出
type Output struct{

}
/*--一笔交易就是包括一个多个输入和一个或多个输出打包成成一个Transaction,然后被挖矿挖出加入到区块链中---*/
type Transaction struct{
	ins   []Input
	outs  []Output
}