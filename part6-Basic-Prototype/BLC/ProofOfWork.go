package BLC

import "math/big"

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, big.NewInt(1)}
}

// Run
func (p *ProofOfWork) Run() ([]byte, int64) {

	return nil, 0
}
