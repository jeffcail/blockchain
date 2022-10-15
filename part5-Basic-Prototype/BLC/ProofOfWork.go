package BLC

type ProofOfWork struct {
	Block *Block
}

// NewProofOfWork
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{block}
}

// Run
func (p *ProofOfWork) Run() ([]byte, int64) {

	return nil, 0
}
