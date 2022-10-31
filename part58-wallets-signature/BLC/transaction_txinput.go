package BLC

import "bytes"

type TXInput struct {
	TxHash    []byte
	Vout      int
	Signature []byte
	PublicKey []byte
}

// UnLockRipemd160Hash
func (txInput *TXInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {
	publicKey := HashPubKey(txInput.PublicKey)
	return bytes.Compare(publicKey, ripemd160Hash) == 0
}
