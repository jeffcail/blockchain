package BLC

import (
	"fmt"
	"log"
	"math/big"

	"github.com/boltdb/bolt"
	"github.com/c/public-chain.io/common/db"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type BlockChain struct {
	Tip []byte // 最新的区块的hash
	DB  *bolt.DB
}

// CreateBlockChainWithGenesisBlock
func CreateBlockChainWithGenesisBlock() *BlockChain {
	db, err := db.OpenBoltDb(dbName)
	if err != nil {
		log.Fatal(err)
	}
	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				return fmt.Errorf("create bucket err: %v", err)
			}
		}
		if b != nil {
			genesisBlock := CreateGenesisBlock("Genesis block......")
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				return fmt.Errorf("put block into bucket err : %v", err)
			}

			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				return fmt.Errorf("put block hash into bucket err : %v", err)
			}
			blockHash = genesisBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return &BlockChain{
		blockHash, db,
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
	var block *Block
	var currentHash []byte = b.Tip
	for {
		err := b.DB.View(func(tx *bolt.Tx) error {
			bt := tx.Bucket([]byte(blockTableName))
			if b != nil {
				blockBytes := bt.Get(currentHash)
				block = DeserializeBlock(blockBytes)
				fmt.Printf("Height: %d\n", block.Height)
				fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %x\n", block.Data)
				fmt.Printf("Timestamp: %d\n", block.Timestamp)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Printf("Nonce: %d\n", block.Nonce)
			}
			return nil
		})

		fmt.Println()
		if err != nil {
			log.Panic(err)
		}

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

	}
}
