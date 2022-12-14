package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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
func NewSimpleTransaction(from, to string, amount int) *Transaction {

	var txIntputs []*TXInput
	var txOutputs []*TXOutPut

	// 代表消费
	bytes, _ := hex.DecodeString("8bfd87a6641157daeff2b54375e05498731896c3c9c962e5efcd5f6484caf4cd")
	txInput := &TXInput{bytes, 0, from}

	// 消费
	txIntputs = append(txIntputs, txInput)

	// 转账
	txOutput := &TXOutPut{4, to}
	txOutputs = append(txOutputs, txOutput)

	// 找零
	txOutput = &TXOutPut{10 - 4, from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txIntputs, txOutputs}
	tx.HashTransaction()
	return tx
}
