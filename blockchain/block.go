package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	data      []byte
	hash      [32]byte
	prev_hash []byte
	//timestamp []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) CreateGenesis() *Block {
	genesis_block := &Block{[]byte("Genesis"), [32]byte{}, []byte{}}
	genesis_block.CalcHash()
	return genesis_block
}

func (b *Block) CalcHash() {
	info := bytes.Join([][]byte{b.data, b.prev_hash}, []byte{})
	b.hash = sha256.Sum256(info)
}

func createBlock(data string, prev_hash []byte) *Block {
	new_block := &Block{[]byte(data), [32]byte{}, prev_hash}
	new_block.CalcHash()
	return new_block
}

func (chain *Blockchain) addBlock(data string) {
	prev_block := chain.blocks[len(chain.blocks)-1]
	new_block := createBlock(data, prev_block.hash[:])
	chain.blocks = append(chain.blocks, new_block)
}
