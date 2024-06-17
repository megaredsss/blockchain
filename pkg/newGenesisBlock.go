package pkg

// Create newGen
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
