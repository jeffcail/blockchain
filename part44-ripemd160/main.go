package main

import (
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	hash := ripemd160.New()
	hash.Write([]byte("www.baidu.com"))
	bytes := hash.Sum(nil)
	fmt.Printf("%x\n", bytes)
}
