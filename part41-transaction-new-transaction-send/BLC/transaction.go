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

// IsCoinbaseTransaction
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}

// NewCoinBaseTransaction
func NewCoinBaseTransaction(address string) *Transaction {
	txInput := &TXInput{
		TxHash:    []byte{},
		Vout:      -1,
		ScriptSig: "Genesis Data",
	}

	txOutput := &TXOutPut{
		Value:        10,
		ScriptPubKey: address,
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
func NewSimpleTransaction(from, to string, amount int, blockChain *BlockChain) *Transaction {
	//unSpentTx := UnSpentTransactionWithAddress(from)
	//fmt.Printf("unSpentTx： %v", unSpentTx)

	money, spendableUTXODic := blockChain.FindSpendableUTXOS(from, amount)

	var txIntputs []*TXInput
	var txOutputs []*TXOutPut

	// 代表消费
	for txHash, indexArray := range spendableUTXODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			txInput := &TXInput{txHashBytes, index, from}
			// 消费
			txIntputs = append(txIntputs, txInput)
		}
	}

	// 转账
	txOutput := &TXOutPut{int64(amount), to}
	txOutputs = append(txOutputs, txOutput)

	// 找零
	txOutput = &TXOutPut{int64(money) - int64(amount), from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txIntputs, txOutputs}
	tx.HashTransaction()

	return tx
}
