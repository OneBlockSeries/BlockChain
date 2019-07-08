# BlockChain
 write Bitcoin
 写个内存版本的bitcoin，仅供学习用，学习bitcoin中一些重要的技术
 包括POW工作量证明, Merkle tree, 创世快，哈希算法应用，比特地址，比特币特有的输入输出，transaction,签名等

1. 很简单的blockchain
2. 完善block结构体,在block结构体中增加BlockHead{time,prevhash,hash}结构体，增加block可以哈希功能
3. 第一个区块，增加创世块GenesisBlock
4. 增加AddBlockToChain函数
5. 修改AddBlockToChain函数，更简单点
6. 增加Input输入， Output输出， Transaction结构体，代表一次交易。block中存储着很多Transaction,
7. 增加coinbase交易，是个特殊的Transaction,,是挖矿奖励，是币的来源
8. 完善Input结构体， 添加sendcoin(from,to string,mount int )也就是转账功能
9. 添加地址 address 功能，然后全部用地址代替集
10. 完善sendcoin功能，从整个区块链中寻找UTXO来作为输入
11. 加入 pow 功能
12. 加入了merkle tree 功能，哈希整个transaction
13. 实现整体的发送币,
14. 完善了从区块链中查找为花费的transaction
15. 完善解锁，加锁功能
16. 加入签名,检验签名
17. 加入网络功能