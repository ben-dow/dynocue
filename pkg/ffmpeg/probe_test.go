package ffmpeg

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFfprobe(t *testing.T) {
	val, err := Probe("/home/benjamin-dow/bigbuckbunny.mp4")
	assert.NotEmpty(t, *val)
	assert.NoError(t, err)

	assert.True(t, val.HasAudioStream())
	assert.True(t, val.HasVideoStream())
	assert.Equal(t, time.Duration(float64(time.Second)*float64(634.6)), val.Duration())
}

func BenchmarkFfprobe(b *testing.B) {
	for b.Loop() {
		Probe("/home/benjamin-dow/bigbuckbunny.mp4")
	}
}

func BenchmarkParseCodecTypes(b *testing.B) {
	val, _ := Probe("/home/benjamin-dow/bigbuckbunny.mp4")

	codecFuncs := []func() bool{
		val.HasAudioStream,
		val.HasVideoStream,
	}

	for _, f := range codecFuncs {
		b.Run(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), func(b *testing.B) {
			for b.Loop() {
				f()
			}
		})
	}

}
