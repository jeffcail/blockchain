package main

import (
	"fmt"

	"github.com/c/public-chain.io/part49-wallet/BLC"
)

func main() {
	wallet := BLC.NewWallet()
	address := wallet.GetAddress()
	fmt.Printf("%s\n", address)

	validateAddress := BLC.ValidateAddress(address)
	fmt.Printf("%v\n", validateAddress)
}
