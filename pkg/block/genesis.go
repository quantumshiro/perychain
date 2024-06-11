package block

func Genesis() *Block {
	return NewBlock(1, []byte{}, []byte{}, []byte{})
}
