package server

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Address string
	router  *chi.Mux
}

type LimeHandler func(body []byte) ([]byte, error)

func (s *Server) AddHandler(meth string, path string, handler LimeHandler) {
	if s.router == nil {
		s.router = chi.NewRouter()
	}
	s.router.MethodFunc(meth, path, createHttpHandler(handler))
}

func (s *Server) Start(port string) error {
	if s.router == nil {
		return errors.New("error, router is nil")
	}
	return http.ListenAndServe(":"+port, s.router)
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
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	}
}
