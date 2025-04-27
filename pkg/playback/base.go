package playback

import (
	"fmt"
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
	cleanupOnce.Do(func() {
		err = vlc.Release()
	})
	return err
}



type MediaPlayer struct {
	mu      sync.RWMutex
	players map[string]*vlc.Player
}

func NewMediaPlayer() *MediaPlayer {
	return &MediaPlayer{
		mu:      sync.RWMutex{},
		players: make(map[string]*vlc.Player),
	}
}

func (mp *MediaPlayer) PlayAudio(id, path string) error {
	mp.mu.Lock()

	_, ok := mp.players[id]
	if ok {
		mp.mu.Unlock()
		return fmt.Errorf("already playing with id %s", id)
	}

	player, err := vlc.NewPlayer()
	if err != nil {
		mp.mu.Unlock()
		return err
	}

	mp.players[id] = player
	mp.mu.Unlock()

	media, err := player.LoadMediaFromPath(path)
	if err != nil {
		return err
	}

	manager, err := player.EventManager()
	if err != nil {
		return err
	}

	quit := make(chan struct{})

	once := sync.Once{}

	eventCallback := func(event vlc.Event, userData interface{}) {
		once.Do(func() {
			close(quit)
		})
	}

	evIdStop, err := manager.Attach(vlc.MediaPlayerStopped, eventCallback, nil)
	if err != nil {
		return err
	}

	evIdEnd, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		return err
	}

	if err = player.Play(); err != nil {
		return err
	}

	go func() {
		defer func() {
			mp.mu.Lock()
			delete(mp.players, id)
			mp.mu.Unlock()

			manager.Detach(evIdStop, evIdEnd)
			media.Release()
			player.Stop()
			player.Release()

		}()

		<-quit
	}()

	return nil
}
