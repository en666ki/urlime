package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		url     string
		isValid bool
	}{
		{"google.com", true},
		{"www.google.com", true},
		{"https://google.com", true},
		{"https://www.google.com", true},
		{"https://www.google.com/asd", true},
		{"http://www.google.com/search?q=dsfsf&rlz=1C5CHFA_enRU984RU984&oq=dsfsf&aqs=chrome..69i57l2j69i59l2.469j0j7&sourceid=chrome&ie=UTF-8", true},
		{"https://https://", false},
		{"https://https://google.com", false},
		{".com", false},
		{"", false},
	}
	for _, tt := range tests {
		result, _ := IsValidUrl(tt.url, Blacklist{})
		assert.Equal(t, tt.isValid, result)

	}
}

func TestShorten(t *testing.T) {
	tests := []struct {
		url  string
		surl string
	}{
		{url: "google.com", surl: "66666626"},
		{url: "http://google.com", surl: "67773226"},
		{url: "", surl: ""},
		{url: "http:/google.com", surl: "67773266"},
		{url: "http://http://http://", surl: "67773226"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.surl, Shorten(tt.url, raw, 8))
	}
}

func raw(data []byte) string {
	result := make([]byte, 0)
	for _, b := range data {
		result = append(result, fmt.Sprintf("%x", b)[0])
	}
	return string(result)
}
