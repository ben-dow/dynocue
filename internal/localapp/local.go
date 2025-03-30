package localapp

import (
	"dynocue/pkg/model"
	"slices"
)

type LocalDynoCue struct {
	show *model.Show
	evCb func(string, interface{})
}

func NewLocalDynoCue(eventCallback func(string, interface{})) *LocalDynoCue {
	ldc := &LocalDynoCue{
		show: model.NewShow(),
		evCb: eventCallback,
	}
	return ldc
}

func OpenLocalDynoCue(eventCallback func(string, interface{})) *LocalDynoCue {
	return &LocalDynoCue{}
}

func (l *LocalDynoCue) notifyShowUpdate() {
	l.evCb("MODEL_UPDATE", map[string]interface{}{"type": "SHOW", "show": l.show})
}

func (l *LocalDynoCue) SetShowName(name string) {
	l.show.Name = name
	l.notifyShowUpdate()
}

func (l *LocalDynoCue) GetShow() *model.Show {
	return l.show
}

func (l *LocalDynoCue) AddAudioSource() {
	l.show.SourceList.AudioSources = append(l.show.SourceList.AudioSources, model.NewAudioSource())
	l.notifyShowUpdate()
}

func (l *LocalDynoCue) UpdateAudioSourceLabel(id, label string) {
	for idx, as := range l.show.SourceList.AudioSources {
		if as.Id == id {
			as.Label = label
			l.show.SourceList.AudioSources[idx] = as
			l.notifyShowUpdate()
			return
		}
	}
}

func (l *LocalDynoCue) DeleteAudioSource(id string) {
	l.show.SourceList.AudioSources = slices.DeleteFunc(l.show.SourceList.AudioSources, func(s model.AudioSource) bool { return s.Id == id })
	l.notifyShowUpdate()
}
