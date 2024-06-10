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

func (b *Block) toString() string {
	return "Block -\n" +
		"Timestamp: " + string(b.timestamp) + "\n" +
		"Last Hash: " + string(b.lastHash) + "\n" +
		"Hash: " + string(b.hash) + "\n" +
		"Data: " + string(b.data) + "\n"
}
