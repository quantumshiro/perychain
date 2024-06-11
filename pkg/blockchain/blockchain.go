package blockchain

import (
	"github.com/quantumshiro/perychain/pkg/block"
)

type BlockChain struct {
	chain []*block.Block
}
