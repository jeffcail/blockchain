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
		Signature: nil,
		PublicKey: []byte{},
	}

	txOutput := NewTXOutput(10, address)

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
func NewSimpleTransaction(from, to string, amount int, blockChain *BlockChain, txs []*Transaction) *Transaction {
	wallets, _ := NewWallets()
	wallet := wallets.WalletsMap[from]

	money, spendableUTXODic := blockChain.FindSpendableUTXOS(from, amount, txs)

	var txIntputs []*TXInput
	var txOutputs []*TXOutPut

	// 代表消费
	for txHash, indexArray := range spendableUTXODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			txInput := &TXInput{txHashBytes, index, nil, wallet.PublicKey}
			// 消费
			txIntputs = append(txIntputs, txInput)
		}
	}

	// 转账
	txOutput := NewTXOutput(int64(amount), to)
	txOutputs = append(txOutputs, txOutput)

	// 找零
	txOutput = NewTXOutput(int64(money)-int64(amount), from)
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txIntputs, txOutputs}
	tx.HashTransaction()

	return tx
}
