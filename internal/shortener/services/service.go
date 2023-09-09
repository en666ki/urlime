package services

import (
	"crypto/md5"
	"fmt"

	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type UrlService struct {
	interfaces.IUrlRepository
}

func (s *UrlService) StoreShortenUrl(url string) (viewmodels.UrlVM, error) {
	surl := shorten(url)
	err := s.PutUrl(surl, url)
	if err != nil {
		return viewmodels.UrlVM{}, err
	}
	return viewmodels.UrlVM{surl, url}, nil
}

func (s *UrlService) ReadUrl(surl string) (models.Url, error) {
	url, err := s.GetUrl(surl)
	if err != nil {
		return models.Url{}, err
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
