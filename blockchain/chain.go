package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Blockchain struct {
	Blocks []*Block
	ledger [][]byte
}

const difficulty = 2

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesis()}, nil}
}

func (chain *Blockchain) AddBlock(Data string, nonce int64) {
	prev_block := chain.Blocks[len(chain.Blocks)-1]
	new_block := createBlock(Data, prev_block.Hash[:], nonce)
	chain.Blocks = append(chain.Blocks, new_block)
}

func (chain *Blockchain) POW() int64 {
	prev_block := chain.Blocks[len(chain.Blocks)-1]
	nonce := prev_block.Nonce
	for VerifyProof(prev_block.Hash, nonce) == false {
		nonce = nonce + 1
	}
	return nonce
}

func VerifyProof(last_Hash [32]byte, nonce int64) bool {
	str := strconv.FormatInt(nonce, 10)
	puzzle := bytes.Join([][]byte{last_Hash[:], []byte(str)}, []byte{})
	guess_Hash := sha256.Sum256(puzzle)

	if string(guess_Hash[:])[:difficulty] == strings.Repeat("0", difficulty) {
		fmt.Printf(string(guess_Hash[:]))
		return true
	}
	return false
}

func (b *Blockchain) BlockMining(miner string, amount string) {
	transaction := bytes.Join([][]byte{[]byte("0"), []byte(miner), []byte(amount)}, []byte(","))
	b.ledger = append(b.ledger, transaction)

	nonce := b.POW()
	b.AddBlock(strings.Join([]string{"0", miner, amount}, ","), nonce)

}

// func (chain *Blockchain) VerifyChain() bool{
// 	i := 0
// 	for range chain.Blocks{
// 	}
// }
