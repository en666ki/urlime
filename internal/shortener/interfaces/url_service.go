package interfaces

type IUrlService interface {
	StoreShortenUrl(url string) (UrlVM, error)
	ReadUrl(surl string) (Url, error)
}
