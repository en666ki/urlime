package service

import (
	"crypto/md5"
	"fmt"
)

type UrlService struct {
	IUrlRepository
}

func (s *UrlService) StoreShortenUrl(url string) (UrlVM, error) {
	surl := shorten(url)
	err := s.PutUrl(surl, url)
	if err != nil {
		return UrlVM{}, err
	}
	return UrlVM{surl, url}, nil
}

func (s *UrlService) ReadUrl(surl string) (Url, error) {
	url, err := s.GetUrl(surl)
	if err != nil {
		return Url{}, err
	}
	return url, nil
}

func shorten(url string) string {
	sum := md5.Sum([]byte(url))
	var s [md5.Size]byte
	for i, b := range sum {
		s[i] = fmt.Sprintf("%x", b)[0]
	}
	return string(s[:8])
}
