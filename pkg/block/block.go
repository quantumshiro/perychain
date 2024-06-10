package block

type Block struct {
	timestamp int64
	lastHash  []byte
	hash      []byte
	data      []byte
}

func NewBlock(timestamp int64, lastHash []byte, hash []byte, data []byte) *Block {
	return &Block{timestamp, lastHash, hash, data}
}
