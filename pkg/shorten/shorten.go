package shorten

import (
	md5 "crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"unicode/utf8"
)

func Shorten(body []byte) ([]byte, int, error) {
	if !hasNoRunes(string(body)) {
		return []byte{}, http.StatusBadRequest, errors.New("no unicode allowed")
	}

	return []byte(getShorten(string(body))[:8]), http.StatusOK, nil
}

func hasNoRunes(s string) bool {
	return utf8.RuneCountInString(s) == len(s)
}

func md5ToString(hash [md5.Size]byte) string {
	var s [md5.Size]byte
	for i, b := range hash {
		s[i] = fmt.Sprintf("%x", b)[0]
	}
	return string(s[:])
}

func getShorten(url string) string {
	sum := md5.Sum([]byte(url))
	return md5ToString(sum)
}
