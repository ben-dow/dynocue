package localapp

import (
	"dynocue/internal/appdef"
	"dynocue/pkg/model"
	"dynocue/pkg/playback"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

type LocalDynoCue struct {
	path    string
	db      *bbolt.DB
	players *playback.PlayerManager

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
		path:    savePath,
		db:      db,
		players: playback.NewPlayerManager(),
		evCb:    eventCallback,
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
		path:    openPath,
		db:      db,
		players: playback.NewPlayerManager(),
		evCb:    eventCallback,
	}

	return ldc, err
}
