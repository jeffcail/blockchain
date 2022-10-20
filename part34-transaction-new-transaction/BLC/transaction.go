package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// UTXO
type Transaction struct {
	TxHash []byte
	Vins   []*TXInput
	Vouts  []*TXOutPut
}

// NewCoinBaseTransaction
func NewCoinBaseTransaction(address string) *Transaction {
	txInput := &TXInput{
		TxID:      []byte{},
		Vout:      -1,
		ScriptSig: "Genesis Data",
	}

	txOutput := &TXOutPut{
		Value:        10,
		ScriptPublic: address,
	}

	txCoinbase := &Transaction{
		TxHash: []byte{},
		Vins:   []*TXInput{txInput},
		Vouts:  []*TXOutPut{txOutput},
	}
	txCoinbase.HashTransaction()
	return txCoinbase
}

// HashTransaction
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

// NewSimpleTransaction
func NewSimpleTransaction(from, to string, amount int) {

}
