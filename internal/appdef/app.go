package appdef

import "dynocue/pkg/model"

type DynoCueApplication interface {
	GetShowMetadata() (*model.ShowMetadata, error)
	SetShowMetadata(metadata *model.ShowMetadata) error
	SetShowName(string) error


	AddAudioSource(Path string)
	UpdateAudioSourceLabel(id, label string)
	DeleteAudioSource(id string)
}
