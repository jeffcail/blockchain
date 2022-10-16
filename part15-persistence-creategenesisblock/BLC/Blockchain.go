package BLC

import (
	"fmt"
	"log"

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
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			return fmt.Errorf("create bucket err: %v", err)
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

//// AddBlockToBlockChain
//func (b *BlockChain) AddBlockToBlockChain(data string, height int64, preBlockHash []byte) {
//	newBlock := NewBlock(data, height, preBlockHash)
//	b.Blocks = append(b.Blocks, newBlock)
//}
