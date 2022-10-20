package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -address -- 交易数据")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT - 交易明细")
	fmt.Println("\tprintchain -- 输出区块信息")
}

func isValidFlags() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(txs []*Transaction) {

	if dbExists() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlockChain(txs)
}

func (cli *CLI) printChain() {

	if dbExists() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.PrintChain()
}

func (cli *CLI) createGenesisBlockchain(address string) {
	CreateBlockChainWithGenesisBlock(address)
}

func (cli *CLI) Run() {
	isValidFlags()

	sendBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "转账源地址...")
	flagTo := sendBlockCmd.String("to", "", "转账目的地地址...")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额...")

	flagCreateBlockChainWithAddress := createBlockchainCmd.String("address", "", "创建创世区块的地址")

	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagAddBlockData)
		//cli.addBlock([]*Transaction{})

		fmt.Println(*flagFrom)
		fmt.Println(*flagTo)
		fmt.Println(*flagAmount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所以区块的数据")
		cli.printChain()
	}

	if createBlockchainCmd.Parsed() {
		if *flagCreateBlockChainWithAddress == "" {
			fmt.Println("地址不能为空...")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainWithAddress)
	}
}
