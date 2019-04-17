package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodingTransform(t *testing.T) {
	samples := []string{
		"8J+Ygwo",
		"8J+Ygwo=",
		"8J-Ygwo",
		"8J-Ygwo=",
	}

	expected := "8J+Ygwo"

	for _, sample := range samples {
		r := reader{bytes.NewReader([]byte(sample))}

		out := make([]byte, len(sample))
		n, err := r.Read(out)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, len(expected), n, "Should read less or equal bytes to input")
		assert.Equal(t, expected, string(out[:n]), "Should correctly map to raw std encoding")
	}
}

func TestReadingChunks(t *testing.T) {
	sample := "8J-Ygwo="
	expected := "8J+Ygwo"

	for c := 1; c < 3; c++ {
		r := reader{bytes.NewReader([]byte(sample))}

		out := []byte{}
		buf := make([]byte, c)
		n, err := r.Read(buf)

		for err == nil {
			out = append(out, buf[:n]...)
			n, err = r.Read(buf)
		}

		assert.Equal(t, expected, string(out), "Should correctly map to raw std encoding")
	}
}
