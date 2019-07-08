package main

import (
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

func StartBlockChain(genesisAddress []byte) (bl *Blockchain) {

	b := GenesisBlock(genesisAddress)
	return &Blockchain{[]*Block{b}}
}

func (bl *Blockchain) Runing(tr *Transaction) {
	fmt.Printf("blockchain running\n")
	//------加入到区块链-------

	//暂时每个块中只有一个Transaction
	bl.AddBlockToChain([]*Transaction{tr})

	//-------把blocks里的区块打印出来-----------
	for idx, value := range bl.blocks {
		fmt.Printf("%d,block,hash=%x\n", idx, value.bhead.hash)
	}
}

func GenesisBlock(genesisAddress []byte) *Block {

	GenesTransaction := Coinbase(genesisAddress)
	Genesisblock := CreateGenesisBlock([]*Transaction{GenesTransaction}) //没有prevhash 第一个块，当然没有前向块的哈希

	return Genesisblock
}
func (bl *Blockchain) AddBlockToChain(trans []*Transaction) {

	block := CreatBlock(trans, bl.blocks[len(bl.blocks)-1].bhead.hash)
	bl.blocks = append(bl.blocks, block)
}

func (bl *Blockchain) SendCoin(from, to []byte, mount int, w *Wallet) *Transaction {

	//封装成transaction
	var ins []Input
	var outs []Output

	//判断address是否合法
	if IsTrueAddress(from) != true || IsTrueAddress(to) != true {
		fmt.Printf("sendcoin  wrong address\n")
		return nil
	}
	//fmt.Printf("%s SendCoin to %s %d coin\n", from, to, mount)
	//两步走 step1,从区块链中找from 找出所有的未花费的输出，
	//step2，封装输入以找出来的输出填 当然得证明是自己的钱
	//封装输出  不要忘记有找零情况

	enoughOutput, findmount := bl.FindEnoughOutputs(from, mount)
	if findmount < mount {
		fmt.Printf("SEND NOT ENOUGH money,%d\n", enoughOutput)
		return nil
	}

	//1.证明是自己的钱，解锁输出，封装输入，2.封装输出，考虑找零情况
	//build a list of inputs
	for txid, outputs := range enoughOutput {

		for _, out := range outputs {
			input := Input{StringToBytes(txid), out, w.PublicKey, []byte{}}

			ins = append(ins, input)
		}
	}

	//build a list of outputs
	output := Output{mount, []byte{}}
	(&output).Lock(to)
	outs = append(outs, output)
	if findmount > mount {
		outs = append(outs, Output{findmount - mount, from}) //找零情况
	}

	tr := Transaction{nil, ins, outs}
	tr.SetId()
	fmt.Printf("sendcoin from=%s,to=%s,mount=%d\n", from, to, mount)
	return &tr
}
func (bl *Blockchain) FindEnoughOutputs(from []byte, needmount int) (map[string][]int, int) {

	unspendOutputs := make(map[string][]int)
	//从区块链中迭代找出住够多的钱
	unspentTransactions := bl.FindUnspentTransactions(from)
	sum := 0

	//从transactions 找出住够的money,并验证
Work:
	for _, tx := range unspentTransactions {
		txid := tx.Id //[]byte transaction哈希
		for outidx, out := range tx.Outs {
			if out.CanbeUnlockWith(from) && sum < needmount {
				sum += out.Value
				unspendOutputs[BytesToString(txid)] = append(unspendOutputs[BytesToString(txid)], outidx)
				if sum >= needmount {
					break Work
				}
			}
		}
	}

	return unspendOutputs, sum
}
func (bl *Blockchain) FindUnspentTransactions(address []byte) []*Transaction {
	var unspentTranS []*Transaction
	
	spentTxos := make(map[string][]int)

	//从尾到头遍历，，如找到合适的Output ，还要验证是否已经花费掉
	lenth:=len(bl.blocks)-1
	for ;lenth>0;lenth--{

		//循环block里的每一个transaction
		for _,tx:=range bl.blocks[lenth].TransactionS{

			//把txid转换成string
			txid:=BytesToString(tx.Id)

		Outputs:
			//循环 transaction 里的输出 OUTPUT
			for outidx,out:=range tx.Outs{
				//out 是否被消费国
				if spentTxos[txid]!=nil{
					for _,spentout:=range spentTxos[txid]{
						if spentout==outidx{
							continue Outputs
						}
					}
				}

				if out.CanbeUnlockWith(address){
					unspentTranS=append(unspentTranS,tx)
				}
			}
			//循环transaction 里面的INPUT,,找出引用之前的输出
			if tx.isCoinbase()==false{
				for _,in:=range tx.Ins{
					if in.CanUnlockOutput(address){
						intxid:=BytesToString(in.Id)
						spentTxos[intxid]=append(spentTxos[intxid],in.OutId)
					}
				}
			}

		} //transaction 循环

	}
	
	return unspentTranS
}
