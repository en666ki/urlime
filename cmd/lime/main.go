package main

import (
	"github.com/en666ki/urlime/pkg/server"
)

func main() {
	var s server.Server
	s.AddHandler("POST", "/test", func(body []byte) ([]byte, error) {
		return body, nil
	})
	s.Start("8080")
}
