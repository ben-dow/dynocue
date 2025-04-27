package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"go.etcd.io/bbolt"
)

func MarshalKeyValueToBucket[T any](db *bbolt.DB, bucket string, key string, value *T) error {
	return db.Batch(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		valBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return b.Put([]byte(key), valBytes)
	})
}

func UnmarshalKeyValueFromBucket[T any](db *bbolt.DB, bucket string, key string, value *T) error {
	return db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket does not exist")
		}
		v := b.Get([]byte(key))
		if len(v) == 0 {
			return fmt.Errorf("not found")
		}
		return json.Unmarshal(v, value)
	})
}

func UnmarshallAllBucketValues[T any](db *bbolt.DB, bucket string) ([]T, error) {
	out := []T{}

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket does not exist")
		}

		err := b.ForEach(func(k, v []byte) error {
			val := new(T)
			err := json.Unmarshal(v, val)
			if err != nil {
				return err
			}
			out = append(out, *val)
			return nil
		})

		return err
	})

	return out, err
}

func MarshalStructToBucket[T any](db *bbolt.DB, bucket string, in *T) error {
	return db.Batch(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		v := reflect.ValueOf(in)

		if v.Kind() != reflect.Pointer && v.Kind() != reflect.Interface {
			return errors.New("must be pointer or interface of struct or map")
		}

		v = v.Elem()

		if v.Kind() != reflect.Struct {
			return errors.New("must be pointer or interface of struct")
		}

		typeOfS := v.Type()
		for i := range typeOfS.NumField() {
			val_bytes, err := json.Marshal(v.Field(i).Interface())
			if err != nil {
				return err
			}
			err = b.Put([]byte(typeOfS.Field(i).Tag.Get("db")), val_bytes)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func UnmarshalStructFromBucket[T any](db *bbolt.DB, bucket string, out *T) error {
	return db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket does not exist")
		}

		v := reflect.ValueOf(out)

		if v.Kind() != reflect.Pointer && v.Kind() != reflect.Interface {
			return errors.New("must be pointer or interface of struct 1")
		}

		v = v.Elem()

		if v.Kind() != reflect.Struct {
			return errors.New("must be pointer or interface of struct")
		}

		typeOfS := v.Type()
		for i := range typeOfS.NumField() {
			byte_value := b.Get([]byte(typeOfS.Field(i).Tag.Get("db")))
			if len(byte_value) == 0 {
				continue
			}

			err := json.Unmarshal(byte_value, v.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
		}

		return nil

	})
}
