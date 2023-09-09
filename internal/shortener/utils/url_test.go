package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasScheme(t *testing.T) {
	tests := []struct {
		url       string
		hasScheme bool
	}{
		{"google.com", false},
		{"http://google.com", true},
		{"", false},
		{"http:/google.com", false},
		{"http://http://http://", true},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.hasScheme, hasScheme(tt.url))
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		s   string
		url string
		ok bool
	}{
		{"google.com", "http://google.com", true},
		{"http://google.com", "http://google.com", true},
		{"", "", false},
		{"http:/google.com", "", false},
		{"http://http://http://", "http://http://http://", true},
	}
	for _, tt := range tests {
		url, err := Validate(tt.s)
		assert.Equal(t, tt.url, url)
		assert.Equal(t, tt.ok, err != nil)
	}
}
