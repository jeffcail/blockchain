package BLC

import (
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
				fmt.Printf("%s\n", in.ScriptSig)
			}
			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Printf("%d\n", out.Value)
				fmt.Printf("%s\n", out.ScriptPubKey)
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
func (b *BlockChain) UnUTXOs(address string) []*TXOutPut {

	var unUTXOS []*TXOutPut
	spentTXOutputs := make(map[string][]int)

	blockIterator := b.Iterator()
	for {
		block := blockIterator.Next()
		fmt.Println(block)
		fmt.Println()

		// Vins
		for _, tx := range block.Txs {
			if tx.IsCoinbaseTransaction() == false {
				for _, in := range tx.Vins {
					if in.UnLockWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}
				}
			}

			// Vouts
			for index, out := range tx.Vouts {
				if out.UnLockScriptPubKeyWithAddress(address) {
					fmt.Println(out)
					if spentTXOutputs != nil {
						for txHash, indexArray := range spentTXOutputs {

							//if txHash == hex.EncodeToString(tx.TxHash) {

							for _, i := range indexArray {
								if index == i && txHash == hex.EncodeToString(tx.TxHash) {
									continue
								} else {
									unUTXOS = append(unUTXOS, out)
								}
							}

							//}

						}
					} else {
						unUTXOS = append(unUTXOS, out)
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

// MineNewBlock
func (b *BlockChain) MineNewBlock(from, to, amount []string) {
	//fmt.Println(from)
	//fmt.Println(to)
	//fmt.Println(amount)

	a, _ := strconv.Atoi(amount[0])
	tx := NewSimpleTransaction(from[0], to[0], a)
	var txs []*Transaction

	txs = append(txs, tx)

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
	utxos := b.UnUTXOs(address)

	var amount int64

	for _, out := range utxos {
		amount = amount + out.Value
	}

	return amount
}
