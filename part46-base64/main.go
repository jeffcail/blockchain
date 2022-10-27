package main

import (
	"encoding/base64"
	"fmt"
)

var base58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func main() {
	bytes := []byte("blog.caixiaoxin.cn")
	encoded := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(encoded)
	decodeString, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decodeString))
}
