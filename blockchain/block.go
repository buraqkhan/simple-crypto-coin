package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Data      []byte
	Hash      [32]byte
	Prev_hash []byte
	Nonce     int64
	//timestamp []byte
}

func CreateGenesis() *Block {
	genesis_block := &Block{[]byte("Genesis"), [32]byte{}, []byte{}, 0}
	genesis_block.CalcHash()
	return genesis_block
}

func createBlock(Data string, Prev_hash []byte, nonce int64) *Block {
	new_block := &Block{[]byte(Data), [32]byte{}, Prev_hash, nonce}
	new_block.CalcHash()
	return new_block
}

func (b *Block) CalcHash() {
	info := bytes.Join([][]byte{b.Data, b.Prev_hash}, []byte{})
	b.Hash = sha256.Sum256(info)
}
