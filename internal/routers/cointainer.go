package routers

import (
	"database/sql"
	"sync"

	"github.com/en666ki/urlime/internal/db/infrastructures"
	"github.com/en666ki/urlime/internal/shortener/repositories"
	"github.com/en666ki/urlime/internal/shortener/services"
	"github.com/en666ki/urlime/internal/shortener/controllers"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectUrlController() (controllers.UrlController, error)
}

type kernel struct{}

func (kernel *kernel) InjectUrlController() (controllers.UrlController, error) {

	sqlConn, err := sql.Open("postgres", "host=postgres_test port=5432 dbname=local user=local password=local_pwd sslmode=disable")
	if err != nil {
		return controllers.UrlController{}, err
	}
	postgresqlHandler := &infrastructures.PostgresqlHandler{}
	postgresqlHandler.Conn = sqlConn

	urlRepository := &repositories.UrlRepository{postgresqlHandler}
	urlService := &services.UrlService{urlRepository}
	urlController := controllers.UrlController{urlService}

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
