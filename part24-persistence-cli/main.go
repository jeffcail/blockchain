package main

import (
	"github.com/c/public-chain.io/part24-persistence-cli/BLC"
)

func main() {
	blockChain := BLC.CreateBlockChainWithGenesisBlock()

	cli := BLC.CLI{blockChain}
	cli.Run()
}
