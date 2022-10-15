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
