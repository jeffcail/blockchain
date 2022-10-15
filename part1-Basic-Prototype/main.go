package main

import (
	"fmt"

	"github.com/c/public-chain.io/part1-Basic-Prototype/BLC"
)

func main() {
	b := BLC.NewBlock("Genesis Block", 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(b)
}
