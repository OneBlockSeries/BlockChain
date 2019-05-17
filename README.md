# BlockChain
 write Bitcoin
 写个内存版本的bitcoin，仅供学习用，学习bitcoin中一些重要的技术

1. 很简单的blockchain
2. 完善block结构体,在block结构体中增加BlockHead{time,prevhash,hash}结构体，增加block可以哈希功能
3. 第一个区块，增加创世块GenesisBlock
4. 增加AddBlockToChain函数
5. 修改AddBlockToChain函数，更简单点
6. 一切为了交易服务，来分解交易，交易分两部分，钱从那里来：挖矿奖励来，暂不考虑交易手续费。钱到哪里去：转账功能
   而所有的钱又可以简单化为两个结构体，输入和输出，输入是正向的，输出是负向的。输出引用输入