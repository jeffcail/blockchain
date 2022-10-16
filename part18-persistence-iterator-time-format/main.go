package main

import (
	"github.com/c/public-chain.io/part18-persistence-iterator-time-format/BLC"
)

func main() {
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	defer blockChain.DB.Close()

	blockChain.AddBlockToBlockChain("Send 100RMD TO liming")
	blockChain.AddBlockToBlockChain("Send 200RMD TO cangjingkong")
	blockChain.AddBlockToBlockChain("Send 300RMD TO xiaozemaliya")
	blockChain.AddBlockToBlockChain("Send 50RMD TO boduoyejieyi")

	blockChain.PrintChain()
}
