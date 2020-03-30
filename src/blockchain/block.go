package blockchain

import (
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

//定义区块
type Block struct {
	Index int            //区块编号
	Timestamp string        //区块时间戳
	PrevBlockHash string //上一个区块的哈希值
	Hash string         //当前区块哈希值
	Data string         //区块数据
}

//计算哈希值
func CalculateHash(b Block) string {
	bIndex := strconv.Itoa(b.Index) //int类型转string类型

	blockData := bIndex + b.Timestamp + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))

	return hex.EncodeToString(hashInBytes[:])
}

//生成新区块
func CreateNewBlock(prevBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.Data = data
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

//创始区块
func CreateFirstBlock() Block {
	prevBlock := Block{}
	prevBlock.Index = -1
	prevBlock.Hash = ""

	return CreateNewBlock(prevBlock, "The first block.")
}