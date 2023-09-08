package routers

import (
	"database/sql"
	"sync"

	"github.com/en666ki/urlime/internal/db/infrastructures"
	"github.com/en666ki/urlime/internal/url"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectUrlController() (url.UrlController, error)
}

type kernel struct{}

func (kernel *kernel) InjectUrlController() (url.UrlController, error) {

	sqlConn, err := sql.Open("postgres", "host=postgres_test port=5432 dbname=local user=local password=local_pwd sslmode=disable")
	if err != nil {
		return url.UrlController{}, err
	}
	postgresqlHandler := &infrastructures.PostgresqlHandler{}
	postgresqlHandler.Conn = sqlConn

	urlRepository := &url.UrlRepository{postgresqlHandler}
	urlService := &url.UrlService{urlRepository}
	urlController := url.UrlController{urlService}

	return urlController, nil
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
