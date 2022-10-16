package db

import (
	"log"

	"github.com/boltdb/bolt"
)

// OpenBoltDb
func OpenBoltDb(dbName string) (*bolt.DB, error) {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return db, nil
}
