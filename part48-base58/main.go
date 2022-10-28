package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/c/public-chain.io/part48-base58/BLC"
)

func main() {
	bytes := []byte("http://www.baidu.com")

	hasher := sha256.New()
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	fmt.Printf("%x\n", hash)

	bytes58 := BLC.Base58Encoding(hash)
	fmt.Printf("%s\n", bytes58)

	decoding := BLC.Base58Decoding(bytes58)
	fmt.Printf("%x\n", decoding[1:])
}
