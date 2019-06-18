package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type BlockHead struct {
	timestamp int64  //创建时间
	prevhash  []byte //前一个区块的哈希值
	hash      []byte //该block的hash值
	nonce     int64     //pow 用到，神奇数字
	merkleRoot  []byte  //merkle tree root ,也就是所有交易哈希根
}
type Block struct {
	TransactionS []*Transaction //交易
	bhead        BlockHead
}

func CreatBlock(trans []*Transaction, prevh []byte) (bl *Block) {

	head := BlockHead{time.Now().Unix(), prevh, []byte{},0,[]byte{}}

	bl = &Block{
		TransactionS: trans,

		bhead: BlockHead{
			timestamp: head.timestamp,
			prevhash:  head.prevhash[:],
			hash:      head.hash[:]}}

	//bl:=&Block{message,head}
	bl.SetHash()
	
	//加入pow后的改动
	pow:=CreateProofOfWork(bl)
	nonce,hash:=pow.PowRun()
	bl.bhead.hash=hash[:]
	bl.bhead.nonce=nonce

	//

	return bl
}
func (b *Block) SetHash() {

	//h:=sha256(head+message)
	time := []byte(strconv.FormatInt(b.bhead.timestamp, 10))

	transactbyte := []byte{}
	blockbytes := bytes.Join([][]byte{b.bhead.prevhash, transactbyte, time}, []byte{})
	h := sha256.Sum256(blockbytes)
	b.bhead.hash = h[:]
	fmt.Printf("sethash=%x\n", b.bhead.hash)
}
