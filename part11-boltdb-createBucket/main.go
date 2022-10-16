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
	createBucket(db)
}

// createBucket
func createBucket(db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		if b != nil {
			err := b.Put([]byte("l"), []byte("send 100 BTC To 阿强..."))
			if err != nil {
				return fmt.Errorf("put value into bucket err: %v", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
