package interfaces

import (
	"github.com/en666ki/urlime/internal/shortener/models"
)

type IUrlRepository interface {
	PutUrl(surl, url string) error
	GetUrl(surl string) (models.Url, error)
}
