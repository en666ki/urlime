package url

type IUrl interface {
	StoreShortenUrl(url string) (UrlVM, error)
	ReadUrl(surl string) (Url, error)
}

type IUrlRepository interface {
	PutUrl(surl, url string) error
	GetUrl(surl string) (Url, error)
}
