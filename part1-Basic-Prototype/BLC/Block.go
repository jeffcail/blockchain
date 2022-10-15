package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/c/public-chain.io/common/utils"
)

type Block struct {
	Height        int64
	Data          []byte
	Timestamp     int64
	Hash          []byte
	PrevBlockHash []byte
}

// NewBlock
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevBlockHash,
	}
	block.SetHash()
	return block
}

// SetHash
func (b *Block) SetHash() {
	heightBytes := utils.IntToBytes(b.Height)
	fmt.Println(heightBytes)

	timestampBytes := []byte(strconv.FormatInt(b.Timestamp, 2))
	fmt.Println(timestampBytes)

	blockBytes := bytes.Join([][]byte{heightBytes, b.PrevBlockHash, b.Data, timestampBytes}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:32]
}
