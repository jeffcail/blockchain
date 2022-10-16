package main

import (
	"flag"
	"fmt"
)

func main() {
	flagPrintChainCmd := flag.String("printchain", "", "输出所有的区块信息")
	flagInt := flag.Int("number", 1, "输出一个整数")
	flagBool := flag.Bool("open", true, "判断真假")

	flag.Parse()
	fmt.Printf("%s\n", *flagPrintChainCmd)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)

}
