package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("www.baidu.com"))
	hashBytes := hash.Sum(nil)
	fmt.Printf("%x\n", hashBytes)
}
