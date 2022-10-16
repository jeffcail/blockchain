package BLC

type BlockChain struct {
	Blocks []*Block
}

// CreateBlockChainWithGenesisBlock
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis block......")
	return &BlockChain{
		Blocks: []*Block{genesisBlock},
	}
}

// AddBlockToBlockChain
func (b *BlockChain) AddBlockToBlockChain(data string, height int64, preBlockHash []byte) {
	newBlock := NewBlock(data, height, preBlockHash)
	b.Blocks = append(b.Blocks, newBlock)
}
