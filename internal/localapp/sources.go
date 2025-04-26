package localapp

import (
	"dynocue/internal/db"
	"dynocue/pkg/model"
	"dynocue/pkg/util"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
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
	val, err := ffmpeg.Probe(inputPath)
	if err != nil {
		return fmt.Errorf("could not find audio source, %w", err)
	}

	codec := gjson.Get(val, "format.format_name")
	codecStr := codec.String()

	durationRaw := gjson.Get(val, "format.duration")
	durationStr := durationRaw.String()
	durationFloat, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		durationFloat = 0
	}
	duration := time.Duration(durationFloat * float64(time.Second))

	if !slices.Contains(util.AudioCodecs(), codecStr) {
		return fmt.Errorf("codec %s is not a supported audio codec", codecStr)
	}

	sourceId := uuid.NewString()
	storagePath := path.Join(l.audioSourcesPath(), sourceId+"."+storageCodec)

	as := &model.AudioSource{
		Id:          sourceId,
		Label:       label,
		StoragePath: storagePath,
		Duration:    duration,
	}

	err = db.MarshalKeyValueToBucket(l.db, AudioSourcesBucket, sourceId, as)
	if err != nil {
		return fmt.Errorf("could not marshal audio source to database, %w", err)
	}

	err = os.MkdirAll(l.audioSourcesPath(), 0755)
	if err != nil {
		return fmt.Errorf("could not create directory structure for audio source, %w", err)
	}

	err = ffmpeg.Input(inputPath).Output(storagePath, ffmpeg.KwArgs{"c:a": storageCodec}).OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("could not import audio source, %w", err)
	}

	return l.notifySources()
}
