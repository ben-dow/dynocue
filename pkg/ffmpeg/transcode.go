package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
)

func TranscodeAudio(source, output, codec string) error {
	cmd := exec.Command("ffmpeg", "-i", source, "-c:a", codec, "-y", output)
	buf := bytes.NewBuffer(nil)
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = stdErrBuf
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[%s] %w", stdErrBuf.String(), err)
	}
	return nil
}
