package main

import (
	"fmt"
	"github.com/c/public-chain.io/part9-searialize-deserializeBlock/BLC"
)

func main() {
	block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	proofOfWork := BLC.NewProofOfWork(block)
	fmt.Printf("%v\n", proofOfWork.IsValid())

	bytes := block.Serialize()
	fmt.Printf("%v\n", bytes)

	block = BLC.DeserializeBlock(bytes)
	fmt.Printf("%v\n", block)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x", block.Hash)
}
