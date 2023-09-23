package services

import (
	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/result"
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

func (s *UrlService) StoreShortenUrl(url string) result.Result {
	surl := utils.Shorten(url)
	err := s.PutUrl(surl, url)
	if err != nil {
		return result.Result{Data: nil, Code: 500, Message: "can't put " + surl + ": " + err.Error()}
	}
	return result.Result{Data: &viewmodels.UrlVM{surl, url}, Code: 200, Message: ""}
}

func (s *UrlService) ReadUrl(surl string) result.Result {
	url, err := s.GetUrl(surl)
	if err != nil {
		return result.Result{Data: nil, Code: 500, Message: "can't get " + surl + ": " + err.Error()}
	}
	vurl := viewmodels.FromModel(url)
	return result.Result{Data: &vurl, Code: 200, Message: ""}
}
