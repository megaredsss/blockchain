package pkg

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
