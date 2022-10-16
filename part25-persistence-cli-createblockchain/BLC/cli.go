package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BlockChain *BlockChain
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -data DATA - 交易数据")
	fmt.Println("\taddblock -data DATA - 交易数据")
	fmt.Println("\tprintchain -- 输出区块信息")
}

func isValidFlags() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BlockChain.AddBlockToBlockChain("send 100 BTC To junCheng......")
}

func (cli *CLI) printChain() {
	cli.BlockChain.PrintChain()
}

func (cli *CLI) createGenesisBlockchain(data string) {
	fmt.Println(data)
}

func (cli *CLI) Run() {
	isValidFlags()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "http://baidu.com", "交易数据")
	flagCreateBlockChainWithData := createBlockchainCmd.String("data", "Genesis block data.....", "创世区块交易数据")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
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

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所以区块的数据")
		cli.printChain()
	}

	if createBlockchainCmd.Parsed() {
		if *flagCreateBlockChainWithData == "" {
			fmt.Println("交易数据不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainWithData)
	}
}
