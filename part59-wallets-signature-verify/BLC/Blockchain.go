package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/c/public-chain.io/common/db"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type BlockChain struct {
	Tip []byte // 最新的区块的hash
	DB  *bolt.DB
}

func (b *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{b.Tip, b.DB}
}

func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateBlockChainWithGenesisBlock
func CreateBlockChainWithGenesisBlock(address string) *BlockChain {
	if dbExists() {
		fmt.Println("创世区块已经存在.......")
		os.Exit(1)
	}

	fmt.Println("正在创建创世区块......")

	db, err := db.OpenBoltDb(dbName)
	if err != nil {
		log.Fatal(err)
	}

	var genesisHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			return fmt.Errorf("create bucket err: %v", err)
		}
		if b != nil {
			txCoinbase := NewCoinBaseTransaction(address)
			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				return fmt.Errorf("put block into bucket err : %v", err)
			}

			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				return fmt.Errorf("put block hash into bucket err : %v", err)
			}
			genesisHash = genesisBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return &BlockChain{
		Tip: genesisHash,
		DB:  db,
	}
}

// AddBlockToBlockChain
func (b *BlockChain) AddBlockToBlockChain(txs []*Transaction) {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		bt := tx.Bucket([]byte(blockTableName))
		if bt != nil {
			// 先获取最新区块
			blockBytes := bt.Get(b.Tip)
			block := DeserializeBlock(blockBytes)

			newBlock := NewBlock(txs, block.Height+1, block.Hash)
			err := bt.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				return fmt.Errorf("err: %v", err)
			}
			err = bt.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				return fmt.Errorf("err: %v", err)
			}
			b.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// PrintChain
func (b *BlockChain) PrintChain() {
	blockChainIterator := b.Iterator()
	for {
		block := blockChainIterator.Next()

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println("Txs:")
		for _, tx := range block.Txs {
			fmt.Printf("%x\n", tx.TxHash)
			fmt.Println("Vins:")
			for _, in := range tx.Vins {
				fmt.Printf("%x\n", in.TxHash)
				fmt.Printf("%d\n", in.Vout)
				fmt.Printf("%s\n", in.PublicKey)
			}
			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Printf("%d\n", out.Value)
				fmt.Printf("%s\n", out.Ripemd160Hash)
			}
		}

		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

// BlockchainObject
func BlockchainObject() *BlockChain {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &BlockChain{tip, db}
}

// UnUTXOs
func (b *BlockChain) UnUTXOs(address string, txs []*Transaction) []*UTXO {
	var unUTXOS []*UTXO
	spentTXOutputs := make(map[string][]int)

	for _, tx := range txs {
		if tx.IsCoinbaseTransaction() == false {
			for _, in := range tx.Vins {
				publicKeyHash := Base58Decoding([]byte(address))
				ripemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
				if in.UnLockRipemd160Hash(ripemd160Hash) {
					key := hex.EncodeToString(in.TxHash)
					spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
				}
			}
		}
	}

	for _, tx := range txs {
	work1:
		for index, out := range tx.Vouts {
			if out.UnLockScriptPubKeyWithAddress(address) {
				if len(spentTXOutputs) == 0 {
					utxo := &UTXO{tx.TxHash, index, out}
					unUTXOS = append(unUTXOS, utxo)
				} else {
					for hash, indexArray := range spentTXOutputs {
						txHashStr := hex.EncodeToString(tx.TxHash)
						if hash == txHashStr {

							var isUnSpentUTXO bool
							for _, outIndex := range indexArray {
								if index == outIndex {
									isUnSpentUTXO = true
									continue work1
								}
								if isUnSpentUTXO == false {
									utxo := &UTXO{tx.TxHash, index, out}
									unUTXOS = append(unUTXOS, utxo)
								}
							}
						} else {
							utxo := &UTXO{tx.TxHash, index, out}
							unUTXOS = append(unUTXOS, utxo)
						}
					}
				}

			}
		}
	}

	blockIterator := b.Iterator()
	for {
		block := blockIterator.Next()
		fmt.Println(block)
		fmt.Println()

		// Vins
		for i := len(block.Txs) - 1; i >= 0; i-- {
			tx := block.Txs[i]
			if tx.IsCoinbaseTransaction() == false {
				for _, in := range tx.Vins {
					publicKeyHash := Base58Decoding([]byte(address))
					ripemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
					if in.UnLockRipemd160Hash(ripemd160Hash) {
						key := hex.EncodeToString(in.TxHash)
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}
				}
			}

			// Vouts
		work:
			for index, out := range tx.Vouts {
				if out.UnLockScriptPubKeyWithAddress(address) {
					fmt.Println(out)
					if spentTXOutputs != nil {
						if len(spentTXOutputs) != 0 {
							var isSpentUTXO bool
							for txHash, indexArray := range spentTXOutputs {
								for _, i := range indexArray {
									if index == i && txHash == hex.EncodeToString(tx.TxHash) {
										isSpentUTXO = true
										continue work
									}
								}
							}
							if isSpentUTXO == false {
								utxo := &UTXO{tx.TxHash, index, out}
								unUTXOS = append(unUTXOS, utxo)
							}
						} else {
							utxo := &UTXO{tx.TxHash, index, out}
							unUTXOS = append(unUTXOS, utxo)
						}
					}
				}
			}
		}

		fmt.Println(spentTXOutputs)
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}

	return unUTXOS
}

// FindSpendableUTXOS
func (b *BlockChain) FindSpendableUTXOS(from string, amount int, txs []*Transaction) (int64, map[string][]int) {
	utxos := b.UnUTXOs(from, txs)
	spendAbleUTXO := make(map[string][]int)

	var value int64
	for _, utxo := range utxos {
		value = value + utxo.Output.Value

		hash := hex.EncodeToString(utxo.TxHash)
		spendAbleUTXO[hash] = append(spendAbleUTXO[hash], utxo.Index)

		if value >= int64(amount) {
			break
		}
	}
	if value < int64(amount) {
		fmt.Printf("%s 's found is not enough\n'", from)
		os.Exit(1)
	}
	return value, spendAbleUTXO
}

// MineNewBlock
func (b *BlockChain) MineNewBlock(from, to, amount []string) {
	var txs []*Transaction
	for index, address := range from {
		a, _ := strconv.Atoi(amount[index])
		tx := NewSimpleTransaction(address, to[index], a, b, txs)
		txs = append(txs, tx)
	}

	var block *Block
	b.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			hash := b.Get([]byte("l"))
			blockBytes := b.Get(hash)
			block = DeserializeBlock(blockBytes)
		}
		return nil
	})

	// 进行签名验证
	for _, tx := range txs {
		if b.VerifyTransaction(tx) != true {
			log.Panic("Verify signatue failed...")
		}
	}

	block = NewBlock(txs, block.Height+1, block.Hash)

	b.DB.Update(func(tx *bolt.Tx) error {

		bx := tx.Bucket([]byte(blockTableName))
		if bx != nil {
			bx.Put(block.Hash, block.Serialize())
			bx.Put([]byte("l"), block.Hash)
			b.Tip = block.Hash
		}

		return nil
	})
}

// GetBalance
func (b *BlockChain) GetBalance(address string) int64 {
	utxos := b.UnUTXOs(address, []*Transaction{})

	var amount int64

	for _, utxo := range utxos {
		amount = amount + utxo.Output.Value
	}

	return amount
}

// SignTransaction
func (b *BlockChain) SignTransaction(tx *Transaction, privateKey ecdsa.PrivateKey) {
	if tx.IsCoinbaseTransaction() {
		return
	}

	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vins {
		prevTX, err := b.FindTransaction(vin.TxHash)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.TxHash)] = prevTX
	}
	tx.Sign(privateKey, prevTXs)
}

// FindTransaction
func (b *BlockChain) FindTransaction(ID []byte) (Transaction, error) {
	bci := b.Iterator()

	for {
		block := bci.Next()
		for _, tx := range block.Txs {
			if bytes.Compare(tx.TxHash, ID) == 0 {
				return *tx, nil
			}
		}

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
	return Transaction{}, nil
}

// VerifyTransaction
func (b *BlockChain) VerifyTransaction(tx *Transaction) bool {
	prevTXs := make(map[string]Transaction)
	for _, vin := range tx.Vins {
		prexTX, err := b.FindTransaction(vin.TxHash)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prexTX.TxHash)] = prexTX
	}
	return tx.Verify(prevTXs)
}
