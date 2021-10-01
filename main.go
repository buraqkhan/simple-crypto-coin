package main

import "blockchainapp/blockchain"

func main() {
	var b = blockchain.Block{[]byte("Test"), [32]byte{}, []byte{}}
	var d = b.CreateGenesis()
	crea
}
