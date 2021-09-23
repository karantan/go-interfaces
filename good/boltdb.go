package good

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

// Define interfaces where they are used and limit them to the bare minimum that is
// required for our code to work.
type DataSource interface {
	View(func(*bolt.Tx) error) error
	Update(func(*bolt.Tx) error) error
}

type Database struct {
	*bolt.DB
}

// GetDatabase returns existing bolt database or creates a new one
func GetDatabase(filename string, readonly bool) (*Database, error) {
	db, err := bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: readonly})
	if err != nil {
		return &Database{}, err
	}
	return &Database{db}, err
}

// Get function retrieves the `key` from the `bucket` using `db` DataSource. Tecnically
// it doesn't care if this is bolt.DB or e.g. badger DB, as long as it has `View` method
// defined as `View(func(*bolt.Tx) error) error` and it uses buckets.
func Get(db DataSource, bucket, key string) (string, error) {
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

// Put function stores the `key` and `value` to the `bucket` using `db` DataSource.
// If the bucket doesn't exist we will create a new bucket.
func Put(db DataSource, bucket, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), []byte(value))
		return err
	})
}
