package shorten

import (
	md5 "crypto/md5"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShorten(t *testing.T) {
	url := "http://google.com"
	referenceSum := md5.Sum([]byte(url))
	reference := md5ToString(referenceSum)
	assert.Equal(t, reference, getShorten(url))
}

func TestMd5ToString(t *testing.T) {
	url := "http://google.com"
	referenceSum := md5.Sum([]byte(url))
	var reference [md5.Size]byte
	for i, b := range referenceSum {
		reference[i] = fmt.Sprintf("%x", b)[0]
	}
	assert.Equal(t, string(reference[:]), md5ToString(referenceSum))
}

func TestHasNoRunes(t *testing.T) {
	tests := []struct {
		givenString string
		hasNoRunes  bool
	}{
		{
			"Hello, world!", true,
		},
		{
			"สวัสดี", false,
		},
		{
			"", true,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.hasNoRunes, hasNoRunes(tt.givenString), tt.givenString)
	}
}
