package routers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"sync"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/db/infrastructures/postgresql"
	"github.com/en666ki/urlime/internal/shortener/controllers"
	"github.com/en666ki/urlime/internal/shortener/repositories"
	"github.com/en666ki/urlime/internal/shortener/services"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectShortenerController(cfg *config.Config, log *slog.Logger) (controllers.UrlController, error)
}

type kernel struct{}

func (kernel *kernel) InjectShortenerController(cfg *config.Config, log *slog.Logger) (controllers.UrlController, error) {
	sqlConn, err := sql.Open(cfg.DB.Driver, fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.SslMode))
	if err != nil {
		return controllers.UrlController{}, err
	}
	postgresqlHandler := postgresql.New(sqlConn)

	urlRepository := repositories.New(postgresqlHandler, cfg, log)
	urlService := services.New(urlRepository, cfg, log)
	urlController := controllers.New(urlService, cfg, log)
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
