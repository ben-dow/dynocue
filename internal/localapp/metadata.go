package localapp

import (
	"dynocue/pkg/model"
	"dynocue/pkg/util"
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

const (
	MetadataBucketName string = "metadata"
)

func (l *LocalDynoCue) SetShowMetadata(metadata *model.ShowMetadata) error {
	err := l.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MetadataBucketName))
		if err != nil {
			return err
		}

		err = util.EncodedStructFields(metadata, "db", func(key string, value []byte) error {
			return b.Put([]byte(key), value)
		})

		return err
	})
	if err != nil {
		return err
	}

	l.notify_update("METADATA", metadata)
	return nil
}

func (l *LocalDynoCue) GetShowMetadata() (*model.ShowMetadata, error) {
	res := &model.ShowMetadata{}
	err := l.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(MetadataBucketName))
		if b == nil {
			return errors.New("metadata bucket does not exist yet")
		}

		err := util.DecodeStructFields(res, "db", func(key string) ([]byte, error) {
			return b.Get([]byte(key)), nil
		})

		return err
	})
	return res, err
}

func (l *LocalDynoCue) SetShowName(n string) error {
	res := &model.ShowMetadata{}
	err := l.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MetadataBucketName))
		if err != nil {
			return err
		}

		nBytes, err := json.Marshal(n)
		if err != nil {
			return err
		}

		err = b.Put([]byte("name"), nBytes)
		if err != nil {
			return err
		}

		err = util.DecodeStructFields(res, "db", func(key string) ([]byte, error) {
			return b.Get([]byte(key)), nil
		})

		return err
	})

	if err != nil {
		return err
	}

	l.notify_update("METADATA", res)
	return nil
}
