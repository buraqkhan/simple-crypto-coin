package main

import (
	"fmt"
 	"blockchainapp/blockchain"
)


func main() {
	a := blockchain.InitBlockChain()
	//a.AddBlock("This is block 2")

	for _, block := range a.Blocks {
		fmt.Printf("prev hash: %x\n", block.Prev_hash)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("data: %s\n", block.Data)
	}
}
