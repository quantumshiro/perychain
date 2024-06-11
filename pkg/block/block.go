package block

import (
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp int64
	LastHash  []byte
	Hash      []byte
	Data      []byte
}

func NewBlock(timestamp int64, lastHash []byte, hash []byte, data []byte) *Block {
	return &Block{timestamp, lastHash, hash, data}
}

func (b *Block) String() string {
	return "Block -\n" +
		"Timestamp: " + string(rune(b.Timestamp)) + "\n" +
		"Last Hash: " + string(b.LastHash) + "\n" +
		"Hash: " + string(b.Hash) + "\n" +
		"Data: " + string(b.Data) + "\n"
}

func MineBlock(lastBlock *Block, data []byte) *Block {
	timestamp := time.Now().Unix()
	var lastHash = lastBlock.Hash
	var hash = []byte("fixme-hash")

	return NewBlock(timestamp, lastHash, hash, data)
}

func hash(timestamp int64, lastHash []byte, data []byte) []byte {
	return sha256.New().Sum([]byte(string(rune(timestamp)) + string(lastHash) + string(data)))
}

func (b *Block) BlockHash() []byte {
	return hash(b.Timestamp, b.LastHash, b.Data)
}
