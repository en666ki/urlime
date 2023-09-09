package routers

import (
	"sync"

	"github.com/en666ki/urlime/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type IShortenRouter interface {
	InitRouter(cfg *config.Config) (*chi.Mux, error)
}

type router struct{}

func (router *router) InitRouter(cfg *config.Config) (*chi.Mux, error) {
	urlController, err := ServiceContainer().InjectUrlController()
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/shorten/{urlParam}", urlController.Shorten)
	r.HandleFunc("/unshort/{surlParam}", urlController.Unshort)
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
