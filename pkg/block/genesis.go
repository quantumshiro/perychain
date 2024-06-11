package block

func (b *Block) Genesis() *Block {
	return NewBlock(1, []byte{}, []byte{}, []byte{})
}
