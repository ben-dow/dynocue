package app

import "dynocue/pkg/model"

type DynoCueApplication interface {
	GetShow() *model.Show
	SetShowName(string)
}
