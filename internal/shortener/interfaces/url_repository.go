package interfaces

type IUrlRepository interface {
	PutUrl(surl, url string) error
	GetUrl(surl string) (Url, error)
}
