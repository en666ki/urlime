package server

import (
	"io"
	"log"
	"net/http"
)

type LimeHandler func(body []byte) ([]byte, error)

func RunHandler(handler LimeHandler, path string, port string) {
	go func() {
		httpHandler := createHttpHandler(handler)

		http.HandleFunc(path, httpHandler)

		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()
}

func createHttpHandler(handler LimeHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to parse HTTP request: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		response, err := handler(body)
		if err != nil {
			log.Printf("Failed to process HTTP request: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		_, err = w.Write(response)
		if err != nil {
			log.Printf("Failed to write HTTP response: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
