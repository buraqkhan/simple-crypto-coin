package main

import (
	"blockchainapp/blockchain"
	"fmt"
)

func main() {
	a := blockchain.InitBlockChain()
	a.BlockMining("Buraq", "1")

	for _, block := range a.Blocks {
		fmt.Printf("prev hash: %x\n", block.Prev_hash)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Println()
	}
}
