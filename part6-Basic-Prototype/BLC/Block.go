package BLC

import (
	"time"
)

type Block struct {
	Height        int64
	Data          []byte
	Timestamp     int64
	Hash          []byte
	PrevBlockHash []byte
	Nonce         int64
}

// NewBlock
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
	}
	// 工作量证明
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// CreateGenesisBlock
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
