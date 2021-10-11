package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strings"
	"strconv"
)

const difficulty = 4

type Block struct {
	Data      []byte
	Hash      [32]byte
	Prev_hash []byte
	Nonce int64
	//timestamp []byte
}

type Blockchain struct {
	Blocks []*Block
}

func CreateGenesis() *Block {
	genesis_block := &Block{[]byte("Genesis"), [32]byte{}, []byte{}, 0}
	genesis_block.CalcHash()
	return genesis_block
}

func (b *Block) CalcHash() {
	info := bytes.Join([][]byte{b.Data, b.Prev_hash}, []byte{})
	b.Hash = sha256.Sum256(info)
}

func createBlock(Data string, Prev_hash []byte, nonce int64) *Block {
	new_block := &Block{[]byte(Data), [32]byte{}, Prev_hash, nonce}
	new_block.CalcHash()
	return new_block
}

func (chain *Blockchain) AddBlock(Data string, nonce int64) {
	prev_block := chain.Blocks[len(chain.Blocks)-1]
	new_block := createBlock(Data, prev_block.Hash[:], nonce)
	chain.Blocks = append(chain.Blocks, new_block)
}

func InitBlockChain() *Blockchain{
	return &Blockchain{[]*Block{CreateGenesis()}}
}

func (chain *Blockchain) POW() int64{
	prev_block := chain.Blocks[len(chain.Blocks)-1]
	nonce := prev_block.Nonce
	for VerifyProof(prev_block.Hash, nonce) == false{
		nonce = nonce + 1
	}
	return nonce
}

func VerifyProof(last_Hash [32]byte, nonce int64) bool{
	str := strconv.FormatInt(nonce, 10)
	puzzle := bytes.Join([][]byte{last_Hash[:], []byte(str)}, []byte{})
	guess_Hash := sha256.Sum256(puzzle)

	return string(guess_Hash[:])[:4] == strings.Repeat("0", difficulty)
}

