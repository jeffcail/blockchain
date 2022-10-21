package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) send(from, to, amount []string) {
	if dbExists() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BlockchainObject()
	defer blockChain.DB.Close()

	blockChain.MineNewBlock(from, to, amount)

}
