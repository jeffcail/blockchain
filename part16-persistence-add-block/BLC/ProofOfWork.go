package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/c/public-chain.io/common/utils"
)

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			utils.IntToBytes(pow.Block.Timestamp),
			utils.IntToBytes(int64(targetBit)),
			utils.IntToBytes(int64(nonce)),
			utils.IntToBytes(pow.Block.Height),
		},
		[]byte{},
	)
	return data
}

// NewProofOfWork
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}

// Run
func (p *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	var hashInt big.Int // 存储新生成的hash
	var hash [32]byte
	for {
		dataBytes := p.prepareData(nonce)
		// 生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if p.Target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1
	}
	return hash[:], int64(nonce)
}
