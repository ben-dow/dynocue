package appdef

import "dynocue/pkg/model"

type DynoCueApplication interface {
	GetShowMetadata() (*model.ShowMetadata, error)
	SetShowMetadata(metadata *model.ShowMetadata) error
	SetShowName(string) error

	GetSources() (*model.Sources, error)
	AddAudioSource(Path string) error
}
