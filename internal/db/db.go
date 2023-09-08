package db

import "errors"

type Url struct {
	Surl string
	Url  string
	db   DB
}

type DB interface {
	PutUrl(url *Url) error
	GetUrl(surl string) (*Url, error)
}

func NewUrl(surl, url string, db DB) *Url {
	return &Url{surl, url, db}
}

func (u *Url) Put() error {
	err := u.db.PutUrl(u)
	if err != nil {
		return err
	}

	return nil
}

func (u *Url) Get() (*Url, error) {
	storedUrl, err := u.db.GetUrl(u.Surl)
	if err != nil {
		return nil, err
	}
	if storedUrl == nil {
		return nil, errors.New("read url is nill")
	}

	u.Surl = storedUrl.Surl
	u.Url = storedUrl.Url

	return storedUrl, nil
}
