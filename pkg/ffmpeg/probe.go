package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/tidwall/gjson"
)

type FileProbe struct {
	raw    string
	parsed gjson.Result
}

func (p *FileProbe) HasAudioStream() bool {
	result := p.parsed.Get(`streams.#(codec_type="audio").codec_name`)
	return result.Exists()
}

func (p *FileProbe) HasVideoStream() bool {
	result := p.parsed.Get(`streams.#(codec_type="video").codec_name`)
	return result.Exists()
}

func (p *FileProbe) Duration() time.Duration {
	return time.Duration(p.parsed.Get("format.duration").Float() * float64(time.Second))
}

func Probe(file string) (*FileProbe, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", file)
	buf := bytes.NewBuffer(nil)
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = stdErrBuf
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", stdErrBuf.String(), err)
	}

	out := new(FileProbe)
	out.raw = buf.String()
	out.parsed = gjson.Parse(buf.String())

	return out, nil
}
