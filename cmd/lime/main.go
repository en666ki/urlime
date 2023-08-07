package main

import (
	"log"

	"github.com/en666ki/urlime/pkg/server"
	"github.com/en666ki/urlime/pkg/shorten"
)

func main() {
	var s server.Server
	s.AddHandler("POST", "/shorten", shorten.Shorten)
	err := s.Start("8080")
	if err != nil {
		log.Fatalf("Server was crushed with error: %v", err)
	}
}
