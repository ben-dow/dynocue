package playback

import (
	"sync"

	vlc "github.com/adrg/libvlc-go/v3"
)

var initOnce sync.Once
var cleanupOnce sync.Once

func InitializePlayback() error {
	var err error
	initOnce.Do(func() {
		err = vlc.Init()
	})
	return err
}

func CleanupPlayback() error {
	var err error
	initOnce.Do(func() {
		err = vlc.Release()
	})
	return err
}
