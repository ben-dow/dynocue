package appdef

import "dynocue/pkg/model"

type DynoCueApplication interface {
	GetShow() *model.Show
	SetShowName(string)

	AddAudioSource()
	UpdateAudioSourceLabel(id, label string)
	DeleteAudioSource(id string)
}
