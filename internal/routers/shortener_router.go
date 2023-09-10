package routers

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/en666ki/urlime/internal/config"
	"github.com/go-chi/chi"
)

type IShortenRouter interface {
	InitRouter(cfg *config.Config, log *slog.Logger) (*chi.Mux, error)
}

type router struct{}

func (router *router) InitRouter(cfg *config.Config, log *slog.Logger) (*chi.Mux, error) {
	urlController, err := ServiceContainer().InjectShortenerController(cfg, log)
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()

	r.HandleFunc(fmt.Sprintf("/shorten/{%s}", cfg.Api.Params.Shorten), urlController.Shorten)
	r.HandleFunc(fmt.Sprintf("/unshort/{%s}", cfg.Api.Params.Unshort), urlController.Unshort)
	return r, nil
}

var (
	m          *router
	routerOnce sync.Once
)

func ShortenRouter() IShortenRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
