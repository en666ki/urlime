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
