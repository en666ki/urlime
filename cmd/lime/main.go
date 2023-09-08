package main

import (
	"log"
	"net/http"

	"github.com/en666ki/urlime/internal/routers"
)

func main() {
	shortenRouter, err := routers.ShortenRouter().InitRouter()
	if err != nil {
		log.Println(err)
		return
	}
	http.ListenAndServe(":8080", shortenRouter)
}
