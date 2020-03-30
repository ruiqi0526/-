package blockchain

import "fmt"

//定义区块链结构体
type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func NewBlockChain() *BlockChain {
	firstBlock := CreateFirstBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&firstBlock)

	return &blockChain
}

//传送数据
func (b *BlockChain) SendData(data string) {
	prevBlock := b.Blocks[len(b.Blocks) - 1]
	newBlock := CreateNewBlock(*prevBlock, data)
	b.AppendBlock(&newBlock)
}

//往链上添加区块
func (b *BlockChain) AppendBlock(newBlock *Block) {
	if len(b.Blocks) == 0 {
		b.Blocks = append(b.Blocks, newBlock)
		return
	}
	if IsValid(*newBlock, *b.Blocks[len(b.Blocks)-1]) {
		b.Blocks = append(b.Blocks, newBlock)
	} else {
		fmt.Println("invalid block!")
	}
}

//验证
func IsValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}