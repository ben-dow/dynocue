package model

type Source interface {
	Label() string
	SetLabel() string
}

type sourceBase struct {
	Label string
}

type AudioSource struct {
	sourceBase
}
