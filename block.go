package main

import (
	//"fmt"
	"bytes"
	"crypto/sha256"
	"time"
	"strconv"
)

type BlockHead struct {
	timestamp int64  //创建时间
	prevhash  []byte //前一个区块的哈希值
	hash      []byte //该block的hash值

}
type Block struct {
	message string
	bhead   BlockHead
}

func CreatBlock(mess string, prevh []byte) (bl *Block) {

	head := BlockHead{time.Now().Unix(), prevh, []byte{}}

	bl = &Block{
		message: mess,

		bhead:BlockHead{
			timestamp: head.timestamp,
			prevhash:  head.prevhash[:],
			hash:      head.hash[:]}}

	//bl:=&Block{message,head}
	bl.SetHash()
	return bl
}
func (b *Block) SetHash() {

	//h:=sha256(head+message)
	time := []byte(strconv.FormatInt(b.bhead.timestamp, 10))
	blockbytes := bytes.Join([][]byte{b.bhead.prevhash,[]byte(b.message),time}, []byte{})
	h := sha256.Sum256(blockbytes)
	b.bhead.hash = h[:]
}
