package main

import (
	"fmt"

	"github.com/c/public-chain.io/part47-wallet/BLC"
)

func main() {
	wallet := BLC.NewWallet()
	fmt.Printf("%s\n", wallet.GetAddress())
}
