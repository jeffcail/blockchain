package BLC

func (cli *CLI) createGenesisBlockchain(address string) {
	blockChain := CreateBlockChainWithGenesisBlock(address)
	defer blockChain.DB.Close()
}
