package localapp

import (
	"dynocue/internal/db"
	"dynocue/pkg/model"
)

const (
	MetadataBucketName string = "metadata"
)

func (l *LocalDynoCue) SetShowMetadata(metadata *model.ShowMetadata) error {
	err := db.MarshalToBucket(l.db, MetadataBucketName, metadata)
	if err != nil {
		return err
	}
	return notify_update[model.ShowMetadata](l.evCb, l.db, "METADATA", MetadataBucketName)
}

func (l *LocalDynoCue) GetShowMetadata() (*model.ShowMetadata, error) {
	res := &model.ShowMetadata{}
	err := db.UnmarshalFromBucket(l.db, MetadataBucketName, res)
	return res, err
}

func (l *LocalDynoCue) setShowMetadataKeyValue(key, value string) error {
	err := db.BatchUpdateValue(l.db, MetadataBucketName, []byte(key), value)
	if err != nil {
		return err
	}
	return notify_update[model.ShowMetadata](l.evCb, l.db, "METADATA", MetadataBucketName)
}

func (l *LocalDynoCue) SetShowName(n string) error {
	return l.setShowMetadataKeyValue("name", n)
}
