package pkg

import (
	"blockchain/pkg/models"
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	models.Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
