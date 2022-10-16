package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	"github.com/c/public-chain.io/part14-block-boltdb/BLC"
)

var (
	db  *bolt.DB
	err error
)

func openBoltDb() {
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func updateBoltBucket(block *BLC.Block) {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				return fmt.Errorf("create bucket err %v\n", err)
			}
		}
		err = b.Put([]byte("l"), block.Serialize())
		if err != nil {
			return fmt.Errorf("put data into bucket err: %v", err)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func readBucket() {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			//fmt.Println(blockData)
			//fmt.Printf("%s\n", blockData)
			block := BLC.DeserializeBlock(blockData)
			fmt.Printf("%v\n", block)
		}
		defer db.Close()
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	openBoltDb()
	updateBoltBucket(block)
	readBucket()
}
