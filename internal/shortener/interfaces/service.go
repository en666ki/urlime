package interfaces

import (
	"github.com/en666ki/urlime/internal/shortener/result"
)

type IUrlService interface {
	StoreShortenUrl(url string) result.Result
	ReadUrl(surl string) result.Result
}
