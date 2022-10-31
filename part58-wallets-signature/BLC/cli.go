package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/c/public-chain.io/common/utils"
)

type CLI struct{}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddresslists -- 输出所有钱包地址")
	fmt.Println("\tcreatewallet -- 创建钱包")
	fmt.Println("\tcreateblockchain -address -- 交易数据")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT - 交易明细")
	fmt.Println("\tprintchain -- 输出区块信息")
	fmt.Println("\tgetbalance -address -- 查询指定账号的余额")
}

func isValidFlags() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	isValidFlags()

	addressListsCmd := flag.NewFlagSet("addresslists", flag.ExitOnError)
	createwalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	sendBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "转账源地址...")
	flagTo := sendBlockCmd.String("to", "", "转账目的地地址...")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额...")

	flagCreateBlockChainWithAddress := createBlockchainCmd.String("address", "", "创建创世区块的地址")
	getBalanceWithAddress := getBalanceCmd.String("address", "", "查询指定账号的余额")

	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "addresslists":
		err := addressListsCmd.Parse(os.Args[2:])
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
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createwalletCmd.Parse(os.Args[2:])
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

		from, _ := utils.JsonToArray(*flagFrom)
		to, _ := utils.JsonToArray(*flagTo)

		for index, fromAddress := range from {
			if IsValidateAddress([]byte(fromAddress)) == false || IsValidateAddress([]byte(to[index])) == false {
				fmt.Printf("地址无效......")
				printUsage()
				os.Exit(1)
			}
		}

		amount, _ := utils.JsonToArray(*flagAmount)
		cli.send(from, to, amount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所以区块的数据")
		cli.printChain()
	}

	if addressListsCmd.Parsed() {
		cli.GetAddressLists()
	}

	if createwalletCmd.Parsed() {
		cli.CreateWallet()
	}

	if createBlockchainCmd.Parsed() {
		if IsValidateAddress([]byte(*flagCreateBlockChainWithAddress)) == false {
			fmt.Println("地址无效...")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainWithAddress)
	}

	if getBalanceCmd.Parsed() {
		if IsValidateAddress([]byte(*getBalanceWithAddress)) == false {
			fmt.Println("地址无效...")
			printUsage()
			os.Exit(1)
		}

		cli.getBalance(*getBalanceWithAddress)
	}
}
