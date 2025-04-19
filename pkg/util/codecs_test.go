package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAudioCodecs(t *testing.T) {
	codecs := AudioCodecs()
	assert.Greater(t, len(codecs), 0)
}

func BenchmarkAudioCodecs(b *testing.B) {
	for b.Loop() {
		AudioCodecs()
	}
}
