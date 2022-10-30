package BLC

import (
	"bytes"
	"math/big"
)

var base58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func ReverseByte(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func Base58Encoding(bytes []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(bytes)

	base := big.NewInt(int64(len(base58)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58[mod.Int64()])
	}

	ReverseByte(result)
	for b := range bytes {
		if b == 0x00 {
			result = append([]byte{base58[0]}, result...)
		} else {
			break
		}
	}
	return result
}

func Base58Decoding(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(base58, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decode := result.Bytes()
	decode = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decode...)

	return decode
}
