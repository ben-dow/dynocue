package api

import "dynocue/pkg/model"

type Api interface {
	NewShow(showPath string) error
	OpenShow(showPath string) error
	CloseShow() error

	SetShowMetadata(metadata *model.ShowMetadata) error
	GetShowMetadata() (*model.ShowMetadata, error)
	SetShowName(n string) error

	GetSources() (*model.Sources, error)
	AddAudioSource(inputPath, storageCodec, label string) error
	PlayAudioSource(sourceId, playbackId string) error
}
