package services

import (
	
	"log"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/utils"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type UrlService struct {
	interfaces.IUrlRepository
	cfg *config.Config
}

func New(repository interfaces.IUrlRepository, cfg *config.Config) *UrlService {
	return &UrlService{repository, cfg}
}

func (s *UrlService) StoreShortenUrl(url string) (viewmodels.UrlVM, error) {
	surl := utils.Shorten(url)
	log.Println(url)
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
