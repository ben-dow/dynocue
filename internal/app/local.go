package app

import (
	"dynocue/pkg/model"
)

type LocalDynoCue struct {
	show *model.Show
	evCb func(string, interface{})

	DynoCueLocalShow
}

func NewLocalDynoCue(eventCallback func(string, interface{})) *LocalDynoCue {
	ldc := &LocalDynoCue{
		show: model.NewShow(),
		evCb: eventCallback,
	}
	ldc.DynoCueLocalShow = NewLocalDynoCueShow(ldc.show, ldc.evCb)
	return ldc
}

func OpenLocalDynoCue(eventCallback func(string, interface{})) *LocalDynoCue {
	return &LocalDynoCue{}
}

type DynoCueLocalShow struct {
	show *model.Show
	evCb func(string, interface{})
}

func NewLocalDynoCueShow(show *model.Show, eventCallback func(string, interface{})) DynoCueLocalShow {
	return DynoCueLocalShow{
		show: show,
		evCb: eventCallback,
	}
}

func (s *DynoCueLocalShow) notifyUpdate() {
	s.evCb("MODEL_UPDATE")
}

func (s *DynoCueLocalShow) SetShowName(name string) {
	s.show.Name = name
	s.notifyUpdate()
}

func (s *DynoCueLocalShow) GetShow() *model.Show {
	return s.show
}
