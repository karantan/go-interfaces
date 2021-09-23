package bad

import (
	"fmt"
	"go-interfaces/logger"
	"time"

	bolt "go.etcd.io/bbolt"
)

var log = logger.New("bad-db")

type Database struct {
	*bolt.DB
}

type Databaser interface {
	Get(string, string) (string, error)
	Put(string, string, string) error
}

// GetDatabase returns existing bolt database or creates a new one
func GetDatabase(filename string, readonly bool) (Databaser, error) {
	db, err := bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: readonly})
	if err != nil {
		return &Database{}, err
	}
	return &Database{db}, nil
}

func (db *Database) Get(bucket, key string) (string, error) {
	var data string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			log.Warnf("Bucket %v doesn't exist!", bucket)
			return fmt.Errorf("Bucket %v doesn't exist!", bucket)
		}
		v := b.Get([]byte(key))
		data = string(v)
		return nil
	})
	return data, err
}

func (db *Database) Put(bucket, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), []byte(value))
		return err
	})
}
