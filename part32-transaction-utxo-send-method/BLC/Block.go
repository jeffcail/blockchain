package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Height        int64
	Txs           []*Transaction
	Timestamp     int64
	Hash          []byte
	PrevBlockHash []byte
	Nonce         int64
}

// NewBlock
func NewBlock(txs []*Transaction, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		Txs:           txs,
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

	fmt.Println()
	return block
}

// CreateGenesisBlock
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

// HashTransactions
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

// Serialize
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// DeserializeBlock
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
