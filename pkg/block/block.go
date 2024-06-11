package block

import (
	"time"
)

type Block struct {
	timestamp int64
	lastHash  []byte
	hash      []byte
	data      []byte
}

func NewBlock(timestamp int64, lastHash []byte, hash []byte, data []byte) *Block {
	return &Block{timestamp, lastHash, hash, data}
}

func (b *Block) String() string {
	return "Block -\n" +
		"Timestamp: " + string(rune(b.timestamp)) + "\n" +
		"Last Hash: " + string(b.lastHash) + "\n" +
		"Hash: " + string(b.hash) + "\n" +
		"Data: " + string(b.data) + "\n"
}

func MineBlock(lastBlock *Block, data []byte) *Block {
	timestamp := time.Now().Unix()
	var lastHash = lastBlock.hash
	var hash = []byte("fixme-hash")

	return NewBlock(timestamp, lastHash, hash, data)
}
