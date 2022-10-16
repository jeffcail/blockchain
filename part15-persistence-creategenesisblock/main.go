package main

import (
	"github.com/c/public-chain.io/part15-persistence-creategenesisblock/BLC"
)

func main() {
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	defer blockChain.DB.Close()

	//blockChain.AddBlockToBlockChain(
	//	"Send 100RMD TO liming",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//blockChain.AddBlockToBlockChain(
	//	"Send 200RMD TO cangjingkong",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//blockChain.AddBlockToBlockChain(
	//	"Send 300RMD TO xiaozemaliya",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//blockChain.AddBlockToBlockChain(
	//	"Send 50RMD TO boduoyejieyi",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//fmt.Println(blockChain)
	//fmt.Println(blockChain.Blocks)
}
