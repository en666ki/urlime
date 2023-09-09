package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
	"github.com/go-chi/chi"
)

type UrlController struct {
	interfaces.IUrlService
	cfg *config.Config
}

func New(service interfaces.IUrlService, cfg *config.Config) *UrlController {
	return &UrlController{service, cfg}
}

func (c *UrlController) Shorten(res http.ResponseWriter, req *http.Request) {
	url := chi.URLParam(req, c.cfg.Api.Params.Shorten)
	log.Println(url)

	storedUrl, err := c.StoreShortenUrl(url)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
	json.NewEncoder(res).Encode(viewmodels.UrlVM{storedUrl.Surl, storedUrl.Url})
}

func (c *UrlController) Unshort(res http.ResponseWriter, req *http.Request) {
	surl := chi.URLParam(req, c.cfg.Api.Params.Unshort)
	log.Println(surl)

	url, err := c.ReadUrl(surl)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
	json.NewEncoder(res).Encode(viewmodels.UrlVM{url.Surl, url.Url})
}
