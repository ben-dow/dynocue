package db

import (
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

func BatchUpdateValue(db *bbolt.DB, bucket string, key []byte, value interface{}) error {
	return db.Batch(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		valBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return b.Put(key, valBytes)
	})
}

func UnmarshalFromBucket[T any](db *bbolt.DB, bucket string, out T) error {
	return db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket does not exist")
		}

		return DecodeStructFields(out, "db", func(key []byte) ([]byte, error) {
			return b.Get(key), nil
		})
	})
}

func MarshalToBucket[T any](db *bbolt.DB, bucket string, in T) error {
	return db.Batch(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		return EncodedStructFields(in, "db", func(key []byte, value []byte) error {
			return b.Put(key, value)
		})
	})
}
