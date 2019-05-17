# BlockChain
 write Bitcoin
 写个内存版本的bitcoin，仅供学习用，学习bitcoin中一些重要的技术

1. 很简单的blockchain
2. 完善block结构体,在block结构体中增加BlockHead{time,prevhash,hash}结构体，增加block可以哈希功能
3. 第一个区块，增加创世块GenesisBlock
4. 增加AddBlockToChain函数
5. 修改AddBlockToChain函数，更简单点
6. 增加Input输入， Output输出， Transaction结构体，代表一次交易。block中存储着很多Transaction,
7. 增加coinbase交易，是个特殊的Transaction,,是挖矿奖励，是币的来源
8. 添加sendcoin(from,to string,mount int )也就是转账功能