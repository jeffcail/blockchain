package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	readBucket(db)
	updateBucket(db)
	readBucket(db)
}

// readBucket
func readBucket(db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		v := b.Get([]byte("l"))
		fmt.Printf("this l is : %s\n", v)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func updateBucket(db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b != nil {
			err := b.Put([]byte("l"), []byte("send 2000 BTC To 阿强..."))
			if err != nil {
				return fmt.Errorf("修改数据失败...")
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
