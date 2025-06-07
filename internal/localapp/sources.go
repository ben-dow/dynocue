package localapp

import (
	"dynocue/internal/db"
	"dynocue/pkg/ffmpeg"
	"dynocue/pkg/model"
	"dynocue/pkg/playback"
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

const (
	AudioSourcesBucket string = "AudioSources"
)

func (l *LocalDynoCue) audioSourcesPath() string {
	return path.Join(l.path, "sources", "audio")
}

func (l *LocalDynoCue) GetSources() (*model.Sources, error) {
	as, err := db.UnmarshallAllBucketValues[model.AudioSource](l.db, AudioSourcesBucket)
	if err != nil {
		return nil, err
	}

	return &model.Sources{
		AudioSources: as,
	}, nil
}

func (l *LocalDynoCue) notifySources() error {
	payload, err := l.GetSources()
	if err != nil {
		return fmt.Errorf("could not retrieve show sources, %w", err)
	}

	l.evCb("MODEL_UPDATE", map[string]interface{}{"type": "SOURCES", "payload": payload})
	return nil
}

func (l *LocalDynoCue) AddAudioSource(inputPath, storageCodec, label string) error {
	probeResult, err := ffmpeg.Probe(inputPath)
	if err != nil {
		return fmt.Errorf("could not find audio source, %w", err)
	}

	if !probeResult.HasAudioStream() {
		return fmt.Errorf("%s does not have a supported audio stream", inputPath)
	}

	sourceId := uuid.NewString()
	storagePath := path.Join(l.audioSourcesPath(), sourceId+"."+storageCodec)

	as := &model.AudioSource{
		Id:          sourceId,
		Label:       label,
		StoragePath: storagePath,
		Duration:    probeResult.Duration(),
	}

	err = db.MarshalKeyValueToBucket(l.db, AudioSourcesBucket, sourceId, as)
	if err != nil {
		return fmt.Errorf("could not marshal audio source to database, %w", err)
	}

	err = os.MkdirAll(l.audioSourcesPath(), 0755)
	if err != nil {
		return fmt.Errorf("could not create directory structure for audio source, %w", err)
	}

	err = ffmpeg.TranscodeAudio(inputPath, storagePath, storageCodec)
	if err != nil {
		return fmt.Errorf("could not import audio source, %w", err)
	}

	return l.notifySources()
}

func (l *LocalDynoCue) PlayAudioSource(id string) error {
	as := new(model.AudioSource)
	err := db.UnmarshalKeyValueFromBucket(l.db, AudioSourcesBucket, id, as)
	if err != nil {
		return err
	}

	err = l.players.StartAudioPlayer(id, &playback.PlayerCfg{
		File: as.StoragePath,
	})

	if err != nil {
		return err
	}

	return nil
}
