package main

import (
	"blockchain/pkg"
	"fmt"
)

func main() {
	bc := pkg.NewBlockchain()
	// Test
	bc.AddBlock("Send 1 BTC to Mary")
	bc.AddBlock("Send 2 more BTC to John")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
