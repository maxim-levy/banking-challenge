package models

import (
	"time"

	"github.com/apex/log"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

// StartBoltDB starts the bolt database.
// https://github.com/boltdb/bolt
func StartBoltDB() {
	// Open the ledger.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open("ledger.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.WithError(err).Fatal("failed to start bolt db")
	}

	// Make data buckets.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(AccountsBucketName))
		if err != nil {
			log.WithError(err).Error("failed to create AccountsBucketName")
			return err
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("failed to make bolt db data buckets")
	}
}

// GetDB returns instance of the DB.
// StartBoltDB should have already been called.
func GetDB() *bolt.DB {
	return db
}
