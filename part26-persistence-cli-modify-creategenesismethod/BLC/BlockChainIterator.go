package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (bci *BlockChainIterator) Next() *Block {
	var block *Block

	err := bci.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockBytes := b.Get(bci.CurrentHash)
			block = DeserializeBlock(currentBlockBytes)

			bci.CurrentHash = block.PrevBlockHash
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return block
}
