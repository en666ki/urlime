package services

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/repositories"
	"github.com/en666ki/urlime/internal/shortener/result"
	"github.com/en666ki/urlime/internal/shortener/utils"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type UrlService struct {
	interfaces.IUrlRepository
	cfg *config.Config
	log *slog.Logger
}

func New(repository interfaces.IUrlRepository, cfg *config.Config, log *slog.Logger) *UrlService {
	return &UrlService{repository, cfg, log}
}

func (s *UrlService) StoreShortenUrl(url string) result.Result {
	surl := utils.Shorten(url)
	err := s.PutUrl(surl, url)
	if err != nil {
		s.log.Error(err.Error(), "domain", "service", "func", "StoreShortenUrl", "url", url)
		code := errorToCode(err)
		return result.Result{Data: nil, Code: code, Message: "can't put " + surl + ": " + err.Error()}
	}
	return result.Result{Data: &viewmodels.UrlVM{surl, url}, Code: http.StatusCreated, Message: ""}
}

func (s *UrlService) ReadUrl(surl string) result.Result {
	url, err := s.GetUrl(surl)
	if err != nil {
		s.log.Error(err.Error(), "domain", "service", "func", "ReadUrl", "surl", surl)
		code := errorToCode(err)
		return result.Result{Data: nil, Code: code, Message: "can't get " + surl + ": " + err.Error()}
	}
	vurl := viewmodels.FromModel(url)
	return result.Result{Data: &vurl, Code: http.StatusOK, Message: ""}
}

func errorToCode(err error) int {
	var code int
	switch {
	case errors.Is(err, repositories.ErrorNotFound):
		code = 404
	case errors.Is(err, repositories.ErrorDatabase):
		code = 500
	case errors.Is(err, repositories.ErrorKeyExists):
		code = 409
	default:
		code = 500
	}
	return code
}
