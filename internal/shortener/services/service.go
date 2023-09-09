package services

import (
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/utils"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type UrlService struct {
	interfaces.IUrlRepository
}

func (s *UrlService) StoreShortenUrl(url string) (viewmodels.UrlVM, error) {

	surl := utils.Shorten(url)
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
