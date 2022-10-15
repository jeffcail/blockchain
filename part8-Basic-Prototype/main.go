package main

import (
	"fmt"
	"github.com/c/public-chain.io/part8-Basic-Prototype/BLC"
)

func main() {
	block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	proofOfWork := BLC.NewProofOfWork(block)
	fmt.Printf("%v\n", proofOfWork.IsValid())
	//blockChain := BLC.CreateBlockChainWithGenesisBlock()
	//
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
