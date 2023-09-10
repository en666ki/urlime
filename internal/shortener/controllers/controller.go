package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
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

	storedUrl, err := c.StoreShortenUrl(url)
	if err != nil {
		c.log.Error("Error", "err", err, "url", url)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	c.log.Info("OK", "surl", storedUrl.Surl, "url", storedUrl.Url)
	json.NewEncoder(res).Encode(viewmodels.UrlVM{storedUrl.Surl, storedUrl.Url})
}

func (c *UrlController) Unshort(res http.ResponseWriter, req *http.Request) {
	surl := chi.URLParam(req, c.cfg.Api.Params.Unshort)

	url, err := c.ReadUrl(surl)
	if err != nil {
		c.log.Error("Error", "err", err, "surl", surl)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	c.log.Info("OK", "surl", url.Surl, "url", url.Url)
	json.NewEncoder(res).Encode(viewmodels.UrlVM{url.Surl, url.Url})
}
