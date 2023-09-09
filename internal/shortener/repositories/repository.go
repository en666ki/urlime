package repository

import (
	"fmt"

	"github.com/en666ki/urlime/internal/db"
)

type UrlRepository struct {
	db.IDHandler
}

func (r *UrlRepository) PutUrl(surl, url string) error {
	_, err := r.Execute(fmt.Sprintf("INSERT INTO %s (surl, url) VALUES ('%s', '%s')", "local_urls", surl, url))
	if err != nil {
		return err
	}
	return nil
}

func (r *UrlRepository) GetUrl(surl string) (Url, error) {
	row, err := r.Query(fmt.Sprintf("SELECT * FROM %s WHERE surl = '%s'", "local_urls", surl))
	if err != nil {
		return Url{}, err
	}
	var url Url
	row.Next()
	row.Scan(&url.Id, &url.Surl, &url.Url)

	return url, nil
}
