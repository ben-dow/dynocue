package model

import (
	"time"

	"github.com/google/uuid"
)

type ShowMetadata struct {
	ShowId string `db:"showId"`
	Name   string `db:"name"`
}

func NewShow() *ShowMetadata {
	return &ShowMetadata{
		ShowId: uuid.NewString(),
	}
}

type CueType string

type CueList struct {
	CueListId string `db:"cueListId"`
	Label     string `db:"label"`
	Cues      []Cue  `db:"cues"`
}

func NewCueList() CueList {
	return CueList{
		CueListId: uuid.NewString(),
		Cues:      make([]Cue, 0),
	}
}

type Cue struct {
	CueId    string  `db:"cueId"`
	CueType  CueType `db:"cueType"`
	Label    string  `db:"label"`
	SourceId string  `db:"sourceId"`

	DelayEnabled bool          `db:"delayEnabled"`
	Delay        time.Duration `db:"delay"`

	FollowEnabled bool          `db:"followEnabled"`
	Follow        time.Duration `db:"follow"`
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
