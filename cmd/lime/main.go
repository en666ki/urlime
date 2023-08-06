package main

import (
	"log"

	"github.com/en666ki/urlime/pkg/server"
)

func main() {
	var s server.Server
	s.AddHandler("POST", "/test", func(body []byte) ([]byte, error) {
		return body, nil
	})
	err := s.Start("8080")
	if err != nil {
		log.Fatalf("Server was crushed with error: %v", err)
	}
}
