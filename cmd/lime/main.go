package main

import (
	"log"

	"github.com/en666ki/urlime/internal/page"
	"github.com/en666ki/urlime/internal/server"
	"github.com/en666ki/urlime/internal/shorten"
)

func main() {
	var s server.Server
	s.AddHandler("GET", "/", page.Main)
	s.AddHandler("POST", "/shorten", shorten.Shorten)
	s.AddHandler("POST", "/unshort", shorten.Unshort)
	err := s.Start("8080")
	if err != nil {
		log.Fatalf("Server was crushed with error: %v", err)
	}
}
