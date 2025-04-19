package util

import (
	"bufio"
	"os/exec"
	"strings"
)

func init() {
	AudioCodecs()
}

var audioCodecs []string

func AudioCodecs() []string {
	if len(audioCodecs) > 0 {
		return audioCodecs
	}

	res := exec.Command("ffmpeg", "-codecs", "-hide_banner", "-loglevel", "error")
	resBytes, err := res.Output()
	if err != nil {
		return []string{}
	}

	resStr := string(resBytes)
	scanner := bufio.NewScanner(strings.NewReader(resStr))

	out := []string{}
	skip := false
	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.Trim(txt, " ")
		if strings.Contains(txt, "Codecs:") {
			skip = true
			continue
		}

		if strings.Contains(txt, "-------") {
			skip = false
			continue
		}

		if skip {
			continue
		}

		fields := strings.Fields(txt)
		data := fields[0]
		codec := fields[1]

		codecData := strings.Split(data, "")
		if codecData[2] == "A" {
			out = append(out, codec)
		}
	}

	audioCodecs = out

	return audioCodecs
}
