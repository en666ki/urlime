package shorten

import (
	md5 "crypto/md5"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/en666ki/urlime/pkg/db"
	"github.com/en666ki/urlime/pkg/db/aerospike"
)

func Unshort(r *http.Request) ([]byte, int, error) {
	err := r.ParseForm()
	if err != nil {
		return []byte{}, http.StatusBadRequest, errors.New("bad unshort form")
	}
	surl := r.Form.Get("surl")
	if !hasNoRunes(string(surl)) {
		return []byte{}, http.StatusBadRequest, errors.New("no unicode allowed")
	}

	log.Printf("\"%s\"", surl)

	url, err := readUrl(surl)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}

	log.Printf("Successfully read: %s %s", surl, url)

	return []byte(url), http.StatusFound, nil
}

func Shorten(r *http.Request) ([]byte, int, error) {
	err := r.ParseForm()
	if err != nil {
		return []byte{}, http.StatusBadRequest, errors.New("bad shorten form")
	}

	url := r.Form.Get("url")
	if !hasNoRunes(url) {
		return []byte{}, http.StatusBadRequest, errors.New("no unicode allowed")
	}

	log.Printf("url: \"%s\"", url)

	surl := getShorten(strings.TrimSpace(url))[:8]
	err = writeShorten(surl, url)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}

	log.Printf("surl: \"%s\"", surl)

	return []byte(surl), http.StatusCreated, nil
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

func writeShorten(surl, url string) error {
	urlToWrite := db.NewUrl(surl, url, aerospike.New("aerospike", 3000, "urlime_test", "url"))
	err := urlToWrite.Put()
	if err != nil {
		return err
	}
	return nil
}

func readUrl(surl string) (string, error) {
	url := db.NewUrl(surl, "", aerospike.New("aerospike", 3000, "urlime_test", "url"))
	_, err := url.Get()
	if err != nil {
		return "", err
	}
	return url.Url, nil
}
