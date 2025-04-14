package util

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Field1 string `db:"field1"`
	Field2 int    `db:"field2"`
	Field3 byte   `db:"field3"`
	Field4 bool   `db:"field4"`
}

func TestEncodeDecodeStruct(t *testing.T) {
	ts := &testStruct{
		Field1: "test",
		Field2: 100,
		Field3: byte(5),
		Field4: true,
	}

	ts_decoded := &testStruct{}

	store := map[string][]byte{}

	err := EncodedStructFields(ts, "db", func(key string, value []byte) error {
		store[key] = value
		return nil
	})
	assert.NoError(t, err)

	err = DecodeStructFields(ts_decoded, "db", func(key string) ([]byte, error) {
		return store[key], nil
	})
	assert.NoError(t, err)

	assert.Equal(t, ts, ts_decoded)
}

func BenchmarkEncode(b *testing.B) {
	ts := &testStruct{
		Field1: "test",
		Field2: 100,
		Field3: byte(5),
		Field4: true,
	}

	for b.Loop() {
		EncodedStructFields(ts, "db", func(key string, value []byte) error { return nil })
	}
}

func BenchmarkDecode(b *testing.B) {

	ts := &testStruct{
		Field1: "test",
		Field2: 100,
		Field3: byte(5),
		Field4: true,
	}
	store := map[string][]byte{}
	EncodedStructFields(ts, "db", func(key string, value []byte) error { store[key] = value; return nil })
	out := &testStruct{}

	for b.Loop() {
		DecodeStructFields(out, "db", func(key string) ([]byte, error) { return store[key], nil })
	}
}

func BenchmarkJson(b *testing.B) {

	ts := &testStruct{
		Field1: "test",
		Field2: 100,
		Field3: byte(5),
		Field4: true,
	}

	for b.Loop() {
		json.Marshal(ts)
	}
}
