package url

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type UrlController struct {
	IUrl
}

func (c *UrlController) Shorten(res http.ResponseWriter, req *http.Request) {
	url := chi.URLParam(req, "urlParam")

	log.Println(req.URL)
	log.Println(url)

	storedUrl, err := c.StoreShortenUrl(url)
	if err != nil {
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
	json.NewEncoder(res).Encode(UrlVM{storedUrl.Surl, storedUrl.Url})
}

func (c *UrlController) Unshort(res http.ResponseWriter, req *http.Request) {
	surl := chi.URLParam(req, "surlParam")

	url, err := c.ReadUrl(surl)
	if err != nil {
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
	json.NewEncoder(res).Encode(UrlVM{url.Surl, url.Url})
}
