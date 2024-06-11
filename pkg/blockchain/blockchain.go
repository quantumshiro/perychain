package blockchain

import (
	"reflect"

	"github.com/quantumshiro/perychain/pkg/block"
)

type BlockChain struct {
	Chain []*block.Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*block.Block{block.Genesis()}}
}

func (bc *BlockChain) AddBlock(data []byte) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := block.MineBlock(lastBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}

// validate the chain
func (bc *BlockChain) Validate(chain *BlockChain) bool {
	if chain.Chain[0].String() != block.Genesis().String() {
		return false
	}

	for i := 1; i < len(chain.Chain); i++ {
		blocks := chain.Chain[i]
		lastBlocks := chain.Chain[i-1]

		if !reflect.DeepEqual(blocks.LastHash, lastBlocks.Hash) || !reflect.DeepEqual(blocks.Hash, blocks.BlockHash()) {
			return false
		}
	}

	return true
}

// replace the chain
func (bc *BlockChain) ReplaceChain(newChain *BlockChain) {
	if len(newChain.Chain) <= len(bc.Chain) {
		return
	}

	if !bc.Validate(newChain) {
		return
	}

	bc.Chain = newChain.Chain
}
