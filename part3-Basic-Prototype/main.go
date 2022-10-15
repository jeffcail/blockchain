package main

import (
	"fmt"

	"github.com/c/public-chain.io/part3-Basic-Prototype/BLC"
)

func main() {
	genesisBlockChain := BLC.CreateBlockChainWithGenesisBlock()
	fmt.Println(genesisBlockChain)
	fmt.Println(genesisBlockChain.Blocks)

	for _, v := range genesisBlockChain.Blocks {
		fmt.Println(v)
	}
}
