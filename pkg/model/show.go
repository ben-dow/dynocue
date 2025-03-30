package model

import (
	"time"

	"github.com/google/uuid"
)

type Show struct {
	ShowId     string     `json:"showId"`
	Name       string     `json:"name"`
	CueList    CueList    `json:"cueLists"`
	SourceList SourceList `json:"sourceList"`
}

func NewShow() *Show {
	return &Show{
		ShowId:     uuid.NewString(),
		CueList:    NewCueList(),
		SourceList: NewSourceList(),
	}
}

type CueType string

type CueList struct {
	CueListId string `json:"cueListId"`
	Label     string `json:"label"`
	Cues      []Cue  `json:"cues"`
}

func NewCueList() CueList {
	return CueList{
		CueListId: uuid.NewString(),
		Cues:      make([]Cue, 0),
	}
}

type Cue struct {
	CueId    string  `json:"cueId"`
	CueType  CueType `json:"cueType"`
	Label    string  `json:"label"`
	SourceId string  `json:"sourceId"`

	DelayEnabled bool          `json:"delayEnabled"`
	Delay        time.Duration `json:"delay"`

	FollowEnabled bool          `json:"followEnabled"`
	Follow        time.Duration `json:"follow"`
}

func NewCue() *Cue {
	return &Cue{
		CueId: uuid.NewString(),
	}
}

type SourceList struct {
	AudioSources []AudioSource
}

func NewSourceList() SourceList {
	return SourceList{
		AudioSources: make([]AudioSource, 0),
	}
}

type AudioSource struct {
	Id    string
	Label string
}

func NewAudioSource() AudioSource {
	return AudioSource{Id: uuid.NewString()}
}
