package ffmpeg

import (
	"bufio"
	"os/exec"
	"strings"
)

var (
	audioCodecs   []string
	videoCodecs   []string
	subtitleCodec []string
	dataCodec     []string
	formats       []string
)

func AudioCodecs() []string {
	if len(audioCodecs) > 0 {
		return audioCodecs
	}
	audioCodecs = Codecs("A")
	return audioCodecs
}

func VideoCodecs() []string {
	if len(videoCodecs) > 0 {
		return videoCodecs
	}
	videoCodecs = Codecs("V")
	return videoCodecs
}

func SubtitleCodecs() []string {
	if len(subtitleCodec) > 0 {
		return subtitleCodec
	}
	subtitleCodec = Codecs("S")
	return subtitleCodec
}

func DataCodecs() []string {
	if len(dataCodec) > 0 {
		return dataCodec
	}
	dataCodec = Codecs("D")
	return dataCodec
}

func Codecs(codecType string) []string {
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
		if codecData[2] == codecType {
			out = append(out, codec)
		}
	}

	return out
}

func Formats() []string {
	if len(formats) > 0 {
		return formats
	}

	res := exec.Command("ffmpeg", "-formats", "-hide_banner", "-loglevel", "error")
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
		if strings.Contains(txt, "File formats:") {
			skip = true
			continue
		}

		if strings.Contains(txt, "--") {
			skip = false
			continue
		}

		if skip {
			continue
		}

		fields := strings.Fields(txt)
		if len(fields) > 0 {
			format := fields[1]
			out = append(formats, format)
		}
	}

	formats = out

	return formats
}
