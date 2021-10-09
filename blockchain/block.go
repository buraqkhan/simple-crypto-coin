package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"strconv"
)

const difficulty = 4

type Block struct {
	data      []byte
	hash      [32]byte
	prev_hash []byte
	Nonce int64
	//timestamp []byte
}

type Blockchain struct {
	blocks []*Block
}

func CreateGenesis() *Block {
	genesis_block := &Block{[]byte("Genesis"), [32]byte{}, []byte{}, 0}
	genesis_block.CalcHash()
	return genesis_block
}

func (b *Block) CalcHash() {
	info := bytes.Join([][]byte{b.data, b.prev_hash}, []byte{})
	b.hash = sha256.Sum256(info)
}

func createBlock(data string, prev_hash []byte, nonce int64) *Block {
	new_block := &Block{[]byte(data), [32]byte{}, prev_hash, nonce}
	new_block.CalcHash()
	return new_block
}

func (chain *Blockchain) AddBlock(data string, nonce int64) {
	prev_block := chain.blocks[len(chain.blocks)-1]
	new_block := createBlock(data, prev_block.hash[:], nonce)
	chain.blocks = append(chain.blocks, new_block)
}

func InitBlockChain() *Blockchain{
	return &Blockchain{[]*Block{CreateGenesis()}}
}

func (chain *Blockchain) POW() int64{
	prev_block := chain.blocks[len(chain.blocks)-1]
	nonce := prev_block.Nonce
	for VerifyProof(prev_block.hash, nonce) == false{
		nonce = nonce + 1
	}
	return nonce
}

func VerifyProof(last_hash [32]byte, nonce int64) bool{
	str := strconv.FormatInt(nonce, 10)
	puzzle := bytes.Join([][]byte{last_hash[:], []byte(str)}, []byte{})
	guess_hash := sha256.Sum256(puzzle)

	return string(guess_hash[:])[:4] == strings.Repeat("0", difficulty)
}


func main() {
	a := InitBlockChain()
	//a.AddBlock("This is block 2")

	for _, block := range a.blocks {
		fmt.Printf("prev hash: %x\n", block.prev_hash)
		fmt.Printf("hash: %x\n", block.hash)
		fmt.Printf("data: %s\n", block.data)
	}
}
