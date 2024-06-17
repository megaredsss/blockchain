package pkg

import (
	"blockchain/pkg/models"
	"time"
)

// Create newblock
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{models.Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{}}}
	block.SetHash()
	return block
}
