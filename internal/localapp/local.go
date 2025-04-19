package localapp

import (
	"dynocue/internal/appdef"
	"dynocue/internal/db"
	"dynocue/pkg/model"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
	"golang.org/x/sync/singleflight"
)

type LocalDynoCue struct {
	path string
	db   *bbolt.DB
	evCb func(string, interface{})
	appdef.NoopDynoCueApplication
}

func NewLocalDynoCue(savePath string, eventCallback func(string, interface{})) (*LocalDynoCue, error) {
	var name string
	if !strings.HasSuffix(savePath, ".dq") {
		name = path.Base(savePath)
		savePath = savePath + ".dq"
	} else {
		suffix, _ := strings.CutSuffix(savePath, ".dq")
		name = path.Base(suffix)
	}

	err := os.Mkdir(savePath, 0755)
	if err != nil {
		return nil, err
	}

	db, err := bbolt.Open(path.Join(savePath, "data.db"), 0600, nil)
	if err != nil {
		return nil, err
	}

	ldc := &LocalDynoCue{
		path: savePath,
		db:   db,
		evCb: eventCallback,
	}

	err = ldc.SetShowMetadata(&model.ShowMetadata{
		ShowId: uuid.NewString(),
		Name:   name,
	})

	if err != nil {
		return nil, err
	}

	return ldc, nil
}

func OpenLocalDynoCue(openPath string, eventCallback func(string, interface{})) (*LocalDynoCue, error) {
	db, err := bbolt.Open(path.Join(openPath, "data.db"), 0600, nil)
	if err != nil {
		return nil, err
	}

	ldc := &LocalDynoCue{
		path: openPath,
		db:   db,
		evCb: eventCallback,
	}

	return ldc, err
}

var sf = &singleflight.Group{}

func notify_update[T any](cb func(string, interface{}), database *bbolt.DB, t string, bucket string) error {
	_, err, _ := sf.Do(t+string(bucket), func() (interface{}, error) {
		payload := new(T)
		err := db.UnmarshalFromBucket(database, bucket, payload)
		if err != nil {
			return nil, err
		}
		cb("MODEL_UPDATE", map[string]interface{}{"type": t, "payload": payload})
		return nil, nil
	})
	return err
}
