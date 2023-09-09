package utils

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

func hasScheme(url string) bool {
	return len(strings.Split(url, "://")) >= 2

}

func Validate(s string) (string, error) {
	if !hasScheme(s) {
		s = "http://" + s
	}
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return "", err
	}
	return s, nil
}

func Shorten(url string) string {
	sum := md5.Sum([]byte(url))
	var s [md5.Size]byte
	for i, b := range sum {
		s[i] = fmt.Sprintf("%x", b)[0]
	}
	return string(s[:8])
}
