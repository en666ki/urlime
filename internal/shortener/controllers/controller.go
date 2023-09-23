package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/go-chi/chi"
)

type UrlController struct {
	interfaces.IUrlService
	cfg *config.Config
	log *slog.Logger
}

func New(service interfaces.IUrlService, cfg *config.Config, log *slog.Logger) *UrlController {
	return &UrlController{service, cfg, log}
}

func (c *UrlController) Shorten(res http.ResponseWriter, req *http.Request) {
	url := chi.URLParam(req, c.cfg.Api.Params.Shorten)

	result := c.StoreShortenUrl(url)
	if result.Data == nil {
		c.log.Error(result.Message, "domain", "controller", "func", "Shorten", "url", url)
		res.WriteHeader(result.Code)
		res.Write([]byte(http.StatusText(result.Code)))
		return
	}
	c.log.Info("OK", "surl", result.Data.Surl, "url", result.Data.Url)
	json.NewEncoder(res).Encode(result)
}

func (c *UrlController) Unshort(res http.ResponseWriter, req *http.Request) {
	surl := chi.URLParam(req, c.cfg.Api.Params.Unshort)

	result := c.ReadUrl(surl)
	if result.Data == nil {
		c.log.Error(result.Message, "domain", "controller", "func", "Unshort", "surl", surl)
		res.WriteHeader(result.Code)
		res.Write([]byte(http.StatusText(result.Code)))
		return
	}
	c.log.Info("OK", "surl", result.Data.Surl, "url", result.Data.Url)
	json.NewEncoder(res).Encode(result)
}
