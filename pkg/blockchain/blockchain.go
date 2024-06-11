package blockchain

import (
	"github.com/quantumshiro/perychain/pkg/block"
)

type BlockChain struct {
	Chain []*block.Block
}

func (bc *BlockChain) AddBlock(data []byte) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := block.MineBlock(lastBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}
