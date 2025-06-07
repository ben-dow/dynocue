package playback

import (
	"fmt"
	"sync"
)

type PlayerManager struct {
	mu      sync.RWMutex
	players map[string]*Player
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		players: map[string]*Player{},
	}
}

func (pm *PlayerManager) StartAudioPlayer(id string, cfg *PlayerCfg) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	_, ok := pm.players[id]
	if ok {
		return fmt.Errorf("source with id %s is already playing", id)
	}

	pl, err := NewAudioPlayer(cfg)
	if err != nil {
		return err
	}

	pm.players[id] = pl

	go func() {
		pl.Wait()
		pm.mu.Lock()
		defer pm.mu.Unlock()

		delete(pm.players, id)
	}()

	return nil
}

func (pm *PlayerManager) GetPlayer(id string) (*Player, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	pl, ok := pm.players[id]
	if !ok {
		return nil, fmt.Errorf("player with id %s doesnt exist", id)
	}

	return pl, nil
}

func (pm *PlayerManager) PlayerIds() []string {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	out := make([]string, 0, len(pm.players))

	for key := range pm.players {
		out = append(out, key)
	}

	return out
}
