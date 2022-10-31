package BLC

import "bytes"

type TXOutPut struct {
	Value         int64
	Ripemd160Hash []byte
}

// Lock
func (txOutput *TXOutPut) Lock(address string) {
	publicKeyHash := Base58Decoding([]byte(address))
	txOutput.Ripemd160Hash = publicKeyHash[1 : len(publicKeyHash)-4]

}

// NewTXOutput
func NewTXOutput(value int64, address string) *TXOutPut {
	txOutput := &TXOutPut{Value: value, Ripemd160Hash: nil}

	// 设置 Ripemd160Hash
	txOutput.Lock(address)

	return txOutput
}

// UnLockScriptPubKeyWithAddress
func (txOutput *TXOutPut) UnLockScriptPubKeyWithAddress(address string) bool {
	publicKeyHash := Base58Decoding([]byte(address))
	hash160 := publicKeyHash[1 : len(publicKeyHash)-4]

	return bytes.Compare(txOutput.Ripemd160Hash, hash160) == 0
}
