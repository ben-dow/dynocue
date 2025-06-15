package dynocue

import (
	"dynocue/internal/db"
	"dynocue/pkg/model"
	"fmt"
)

const (
	MetadataBucketName string = "metadata"
)

func (l *LocalDynoCue) SetShowMetadata(metadata *model.ShowMetadata) error {
	err := db.MarshalStructToBucket(l.db, MetadataBucketName, metadata)
	if err != nil {
		return err
	}
	return l.notifyShowMetadata()
}

func (l *LocalDynoCue) GetShowMetadata() (*model.ShowMetadata, error) {
	res := &model.ShowMetadata{}
	err := db.UnmarshalStructFromBucket(l.db, MetadataBucketName, res)
	return res, err
}

func (l *LocalDynoCue) setShowMetadataKeyValue(key, value string) error {
	err := db.MarshalKeyValueToBucket(l.db, MetadataBucketName, key, &value)
	if err != nil {
		return err
	}
	return l.notifyShowMetadata()
}

func (l *LocalDynoCue) SetShowName(n string) error {
	return l.setShowMetadataKeyValue("name", n)
}

func (l *LocalDynoCue) notifyShowMetadata() error {
	payload, err := l.GetShowMetadata()
	if err != nil {
		return fmt.Errorf("could not retrieve show metadata, %w", err)
	}

	l.evCb("MODEL_UPDATE", map[string]interface{}{"type": "METADATA", "payload": payload})
	return nil
}
