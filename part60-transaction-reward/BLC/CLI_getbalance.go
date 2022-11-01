package BLC

import "fmt"

func (cli *CLI) getBalance(address string) {
	fmt.Println("地址: " + address)
	blockChain := BlockchainObject()
	defer blockChain.DB.Close()

	amount := blockChain.GetBalance(address)
	fmt.Printf("%s一共有%d个Token\n", address, amount)
}
