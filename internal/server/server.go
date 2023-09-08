package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Address string
	router  *chi.Mux
}

type LimeHandler func(r *http.Request) ([]byte, int, error)

func (s *Server) AddHandler(meth string, path string, handler LimeHandler) {
	if s.router == nil {
		s.router = chi.NewRouter()
		s.router.Use(middleware.Logger)
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

		response, code, err := handler(r)
		if err != nil {
			if code == 0 {
				code = http.StatusInternalServerError
			}
			log.Printf("Failed to process HTTP request: %v", err)
			http.Error(w, http.StatusText(code), code)
		}

		_, err = w.Write(response)
		if err != nil {
			log.Printf("Failed to write HTTP response: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
