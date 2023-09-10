package interfaces

import (
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type IUrlService interface {
	StoreShortenUrl(url string) (viewmodels.UrlVM, error)
	ReadUrl(surl string) (models.Url, error)
}
