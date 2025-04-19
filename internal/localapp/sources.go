package localapp

import (
	"dynocue/pkg/model"
	"dynocue/pkg/util"
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func (l *LocalDynoCue) GetSources() (*model.Sources, error) {
	return &model.Sources{}, nil
}

func (l *LocalDynoCue) AddAudioSource(inputPath string) error {
	val, err := ffmpeg.Probe(inputPath)
	if err != nil {
		return err
	}

	util.AudioCodecs()

	fmt.Println(val)

	return nil
}
