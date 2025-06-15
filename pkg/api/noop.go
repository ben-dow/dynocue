package api

import "dynocue/pkg/model"

type Noop struct{}

func (n Noop) SetShowMetadata(metadata *model.ShowMetadata) error {
	return nil
}

func (n Noop) GetShowMetadata() (*model.ShowMetadata, error) {
	return nil, nil
}

func (n Noop) SetShowName(string) error {
	return nil
}

func (n Noop) GetSources() (*model.Sources, error) {
	return nil, nil
}

func (n Noop) AddAudioSource(inputPath, storageCodec, label string) error {
	return nil
}

func (n Noop) PlayAudioSource(sourceId, playbackId string) error {
	return nil
}

func (Noop) NewShow(showPath string) error {
	return nil
}

func (Noop) OpenShow(showPath string) error {
	return nil
}

func (Noop) CloseShow() error {
	return nil
}
