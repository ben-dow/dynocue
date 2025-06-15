package playback

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"time"
)

const portRangeLow int = 49152
const portRangeHigh int = 65535

func getPort() int {
	return rand.Intn(portRangeHigh-portRangeLow) + portRangeLow
}

func getPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range length {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

type Player struct {
	cmd *exec.Cmd

	rcAddr     string
	rcPort     string
	rcPassword string
	rcUrl      string
	httpClient http.Client

	closed   chan struct{}
	closeErr error
}

type PlayerCfg struct {
	File      string
	AudioOnly bool
}

func NewAudioPlayer(cfg *PlayerCfg) (*Player, error) {
	addr := "127.0.0.1"
	port := strconv.FormatInt(int64(getPort()), 10)
	password := getPassword(8)
	cmd := exec.CommandContext(context.Background(), "cvlc",
		"-I", "http",
		"--http-host", addr,
		"--http-port", port,
		"--http-password", password,
		"--play-and-exit",
		cfg.File,
	)
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	p := &Player{
		cmd:        cmd,
		rcAddr:     addr,
		rcPort:     port,
		rcPassword: password,
		rcUrl:      fmt.Sprintf("http://%s:%s", addr, port),

		closed: make(chan struct{}),
	}

	go func() {
		p.closeErr = p.cmd.Wait()
		close(p.closed)
	}()

	for _, err := p.Status(); err != nil; _, err = p.Status() {
		time.Sleep(10 * time.Millisecond)
	}

	cmd.Stdout = os.Stdout

	return p, nil
}

type PlayerStatus struct {
	Time          int            `json:"time"`
	Rate          int            `json:"rate"`
	Fullscreen    int            `json:"fullscreen"`
	AudioFilters  map[string]any `json:"audiofilters"`
	Length        int            `json:"length"`
	CurrentPlId   int            `json:"currentplid"`
	Position      int            `json:"position"`
	Equalizer     []any          `json:"equalizer"`
	Random        bool           `json:"random"`
	ApiVersion    int            `json:"apiversion"`
	Version       string         `json:"version"`
	Repeat        bool           `json:"repeat"`
	Loop          bool           `json:"loop"`
	State         string         `json:"state"`
	VideoEffects  map[string]any `json:"videoeffects"`
	Volume        int            `json:"volume"`
	AudioDelay    int            `json:"audiodelay"`
	SubtitleDelay int            `json:"subtitledelay"`
	SeekSeconds   int            `json:"seek_sec"`
}

const (
	pathStatus   string = "status.json"
	pathPlaylist string = "playlist.json"
	pathBrowse   string = "browse.json"
)

func (p *Player) command(path, command string, args ...string) ([]byte, error) {
	select {
	case <-p.closed:
		return nil, fmt.Errorf("closed")
	default:
	}

	cmdStr := ""
	if command != "" {
		cmdStr = "?command=" + command
	}

	if len(args) > 0 {
		if len(args)%2 != 0 {
			return nil, errors.New("args must be provided in sets of 2")
		}

		for pair := range slices.Chunk(args, 2) {
			cmdStr = cmdStr + "&" + pair[0] + "=" + pair[1]
		}
	}

	req, err := http.NewRequest("GET", p.rcUrl+"/requests/"+path+cmdStr, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", p.rcPassword)
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (p *Player) Status() (*PlayerStatus, error) {
	statusBody, err := p.command(pathStatus, "")
	if err != nil {
		return nil, err
	}

	out := &PlayerStatus{}
	err = json.Unmarshal(statusBody, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (p *Player) Play() error {
	_, err := p.command(pathStatus, "command", "pl_play")
	if err != nil {
		return err
	}

	return nil
}
func (p *Player) Pause() error {
	_, err := p.command(pathStatus, "command", "pl_pause")
	if err != nil {
		return err
	}

	return nil
}
func (p *Player) Stop() error {
	_, err := p.command(pathStatus, "command", "pl_stop")
	if err != nil {
		return err
	}

	err = p.Close()
	if err != nil {
		return err
	}

	return nil
}

func (p *Player) Wait() {
	<-p.closed
}

func (p *Player) Close() error {
	select {
	case <-p.closed:
		return nil
	default:
	}

	err := p.cmd.Process.Signal(os.Interrupt)
	if err != nil {
		return err
	}

	<-p.closed
	return p.closeErr
}
