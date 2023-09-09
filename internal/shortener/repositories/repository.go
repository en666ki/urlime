package repositories

import (
	"fmt"
	"log"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/db"
	"github.com/en666ki/urlime/internal/shortener/models"
)

type UrlRepository struct {
	db.IDHandler
	cfg *config.Config
}

func New(handler db.IDHandler, cfg *config.Config) *UrlRepository {
	return &UrlRepository{handler, cfg}
}

func (r *UrlRepository) PutUrl(surl, url string) error {
	_, err := r.Execute(fmt.Sprintf("INSERT INTO %s (surl, url) VALUES ('%s', '%s')", r.cfg.DB.Table, surl, url))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *UrlRepository) GetUrl(surl string) (models.Url, error) {
	row, err := r.Query(fmt.Sprintf("SELECT * FROM %s WHERE surl = '%s'", r.cfg.DB.Table, surl))
	if err != nil {
		return models.Url{}, err
	}
	var url models.Url
	row.Next()
	row.Scan(&url.Id, &url.Surl, &url.Url)

	return url, nil
}
