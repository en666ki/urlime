package controllers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces/mocks"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	urlService := new(mocks.MockUrlService)

	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	urlService.On("StoreShortenUrl", "testurl").Return(viewmodels.UrlVM{"testsurl", "testurl"}, nil)

	urlController := New(urlService, config.MustLoad(), log)
	req := httptest.NewRequest("GET", "http://localhost:8080/shorten/testurl", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()

	r.HandleFunc("/shorten/{url}", urlController.Shorten)

	r.ServeHTTP(w, req)

	expecterResult := viewmodels.UrlVM{}
	expecterResult.Url = "testurl"
	expecterResult.Surl = "testsurl"

	actualResult := viewmodels.UrlVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expecterResult, actualResult)
}

func TestShortenError(t *testing.T) {
	urlService := new(mocks.MockUrlService)

	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	urlService.
		On("StoreShortenUrl", "testurl").
		Return(viewmodels.UrlVM{}, errors.New("oops! gremlins broke Shorten handler!"))

	urlController := New(urlService, config.MustLoad(), log)
	req := httptest.NewRequest("GET", "http://localhost:8080/shorten/testurl", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()

	r.HandleFunc("/shorten/{url}", urlController.Shorten)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestUnshort(t *testing.T) {
	urlService := new(mocks.MockUrlService)

	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	urlService.On("ReadUrl", "testsurl").Return(models.Url{"testsurl", "testurl"}, nil)

	urlController := New(urlService, config.MustLoad(), log)
	req := httptest.NewRequest("GET", "http://localhost:8080/unshort/testsurl", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()

	r.HandleFunc("/unshort/{surl}", urlController.Unshort)

	r.ServeHTTP(w, req)

	expecterResult := viewmodels.UrlVM{}
	expecterResult.Url = "testurl"
	expecterResult.Surl = "testsurl"

	actualResult := viewmodels.UrlVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expecterResult, actualResult)
}

func TestUnshortError(t *testing.T) {
	urlService := new(mocks.MockUrlService)

	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	urlService.
		On("ReadUrl", "testsurl").
		Return(models.Url{}, errors.New("oops! gremlins broke Unshort handler!"))

	urlController := New(urlService, config.MustLoad(), log)
	req := httptest.NewRequest("GET", "http://localhost:8080/unshort/testsurl", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()

	r.HandleFunc("/unshort/{surl}", urlController.Unshort)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
