package main

import (
	"fmt"

	"github.com/c/public-chain.io/part53-wallet/BLC"
)

func main() {
	wallets := BLC.NewWallets()

	wallets.CreateNewWallet()
	wallets.CreateNewWallet()

	fmt.Println(wallets.Wallets)
}
