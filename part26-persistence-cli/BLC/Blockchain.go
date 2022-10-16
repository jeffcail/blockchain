package BLC

import (
	"fmt"
	"log"
	"math/big"
	"os"
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
func CreateBlockChainWithGenesisBlock(data string) {
	if dbExists() {
		fmt.Println("创世区块已经存在.......")
		os.Exit(1)
	}

	fmt.Println("正在创建创世区块......")

	db, err := db.OpenBoltDb(dbName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			return fmt.Errorf("create bucket err: %v", err)
		}
		if b != nil {
			genesisBlock := CreateGenesisBlock(data)
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				return fmt.Errorf("put block into bucket err : %v", err)
			}

			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				return fmt.Errorf("put block hash into bucket err : %v", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// AddBlockToBlockChain
func (b *BlockChain) AddBlockToBlockChain(data string) {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		bt := tx.Bucket([]byte(blockTableName))
		if bt != nil {
			// 先获取最新区块
			blockBytes := bt.Get(b.Tip)
			block := DeserializeBlock(blockBytes)

			newBlock := NewBlock(data, block.Height+1, block.Hash)
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
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}
