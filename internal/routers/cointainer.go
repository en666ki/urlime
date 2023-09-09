package routers

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/db/infrastructures/postgresql"
	"github.com/en666ki/urlime/internal/shortener/controllers"
	"github.com/en666ki/urlime/internal/shortener/repositories"
	"github.com/en666ki/urlime/internal/shortener/services"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectUrlController(cfg *config.Config) (controllers.UrlController, error)
}

type kernel struct{}

func (kernel *kernel) InjectUrlController(cfg *config.Config) (controllers.UrlController, error) {
	log.Printf("open sql connection: host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.SslMode)
	sqlConn, err := sql.Open(cfg.DB.Driver, fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.SslMode))
	if err != nil {
		return controllers.UrlController{}, err
	}
	postgresqlHandler := postgresql.New(sqlConn)

	urlRepository := repositories.New(postgresqlHandler, cfg)
	urlService := services.New(urlRepository, cfg)
	urlController := controllers.New(urlService, cfg)
	return *urlController, nil
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