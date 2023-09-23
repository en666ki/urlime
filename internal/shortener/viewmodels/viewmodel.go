package viewmodels

import "github.com/en666ki/urlime/internal/shortener/models"

type UrlVM struct {
	Surl string
	Url  string
}

func FromModel(url models.Url) UrlVM {
	return UrlVM{Surl: url.Surl, Url: url.Url}
}
