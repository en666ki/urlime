package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/db"
	"github.com/en666ki/urlime/internal/shortener/models"
)

type UrlRepository struct {
	db.IDHandler
	cfg *config.Config
}

var (
	ErrorNotFound  = errors.New("not found")
	ErrorKeyExists = errors.New("key exists")
	ErrorDatabase  = errors.New("database internal error")
)

func New(handler db.IDHandler, cfg *config.Config) *UrlRepository {
	return &UrlRepository{handler, cfg}
}

func (r *UrlRepository) PutUrl(surl, url string) error {
	_, err := r.Execute(fmt.Sprintf("INSERT INTO %s (surl, url) VALUES ('%s', '%s')", r.cfg.DB.Table, surl, url))
	if err != nil {
		e := err.Error()
		fmt.Print(e)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"urls_pkey\"") {
			return ErrorKeyExists
		}
		return ErrorDatabase
	}
	return nil
}

func (r *UrlRepository) GetUrl(surl string) (models.Url, error) {
	row, err := r.Query(fmt.Sprintf("SELECT * FROM %s WHERE surl = '%s'", r.cfg.DB.Table, surl))
	if err != nil {
		return models.Url{}, ErrorDatabase
	}
	var url models.Url
	row.Next()
	err = row.Scan(&url.Surl, &url.Url)
	if err != nil {
		return models.Url{}, ErrorNotFound
	}

	return url, nil
}
