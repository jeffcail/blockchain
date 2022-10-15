package main

import (
	"fmt"

	"github.com/c/public-chain.io/part2-Basic-Prototype/BLC"
)

func main() {
	b := BLC.CreateGenesisBlock("Genesis Block")
	fmt.Println(b)
}
