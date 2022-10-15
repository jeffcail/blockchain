package main

import (
	"fmt"

	"github.com/c/public-chain.io/part5-Basic-Prototype/BLC"
)

func main() {
	blockChain := BLC.CreateBlockChainWithGenesisBlock()

	blockChain.AddBlockToBlockChain(
		"Send 100RMD TO liming",
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlockToBlockChain(
		"Send 200RMD TO cangjingkong",
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlockToBlockChain(
		"Send 300RMD TO xiaozemaliya",
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlockToBlockChain(
		"Send 50RMD TO boduoyejieyi",
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	fmt.Println(blockChain)
	fmt.Println(blockChain.Blocks)

	for _, v := range blockChain.Blocks {
		fmt.Println(v)
	}
}
