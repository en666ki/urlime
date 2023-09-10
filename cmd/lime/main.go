package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/routers"
)

func main() {
	cfg := config.MustLoad()

	log := newLogger(cfg)

	shortenRouter, err := routers.ShortenRouter().InitRouter(cfg, log)
	if err != nil {
		log.Error("%v", err)
		return
	}
	http.ListenAndServe(cfg.Server.Host+cfg.Server.Port, shortenRouter)
}

func newLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
