package main

import (
	"log"
	"net/http"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/routers"
)

func main() {
	cfg := config.MustLoad()

	shortenRouter, err := routers.ShortenRouter().InitRouter(cfg)
	if err != nil {
		log.Println(err)
		return
	}
	http.ListenAndServe(cfg.Server.Host+cfg.Server.Port, shortenRouter)
}
