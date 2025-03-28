package model

import "github.com/google/uuid"

type Show struct {
	ShowId   string       `json:"showId"`
	Metadata ShowMetadata `json:"metadata"`
	CueLists CueList      `json:"cueLists"`
}

type ShowMetadata struct {
	Name string `json:"name"`
}

func NewShow() *Show {
	return &Show{
		ShowId:   uuid.NewString(),
		CueLists: *NewCueList(),
		Metadata: ShowMetadata{},
	}
}
