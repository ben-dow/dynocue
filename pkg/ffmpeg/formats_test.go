package ffmpeg

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodecs(t *testing.T) {
	codecFuncs := []func() []string{
		AudioCodecs,
		VideoCodecs,
		SubtitleCodecs,
		DataCodecs,
	}

	for _, f := range codecFuncs {
		t.Run(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), func(t *testing.T) {
			codecs := f()
			assert.Greater(t, len(codecs), 0)

			codecsCached := f()
			assert.Greater(t, len(codecsCached), 0)
		})
	}
}

func BenchmarkCodecs(b *testing.B) {
	codecFuncs := []func() []string{
		AudioCodecs,
		VideoCodecs,
		SubtitleCodecs,
		DataCodecs,
	}

	for _, f := range codecFuncs {
		b.Run(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), func(b *testing.B) {
			for b.Loop() {
				f()
			}
		})
	}
}

func TestFileFormats(t *testing.T) {
	formats := Formats()
	assert.Greater(t, len(formats), 0)
}

func BenchmarkFileFormats(b *testing.B) {
	for b.Loop() {
		Formats()
	}
}
