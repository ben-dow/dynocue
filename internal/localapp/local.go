package localapp

import (
	"dynocue/internal/appdef"
	"dynocue/pkg/model"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
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
		name = savePath
		savePath = savePath + ".dq"
	} else {
		name, _ = strings.CutSuffix(savePath, ".dq")
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

func (l *LocalDynoCue) notify_update(t string, payload interface{}) {
	l.evCb("MODEL_UPDATE", map[string]interface{}{"type": t, "payload": payload})
}

func (l *LocalDynoCue) SaveShow() {
}
