package controllers

import (
	"encoding/json"
	"net/http/httptest"
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

	urlService.On("StoreShortenUrl", "testurl").Return(viewmodels.UrlVM{"testsurl", "testurl"}, nil)

	urlController := New(urlService, config.MustLoad())
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

func TestUnshort(t *testing.T) {
	urlService := new(mocks.MockUrlService)

	urlService.On("ReadUrl", "testsurl").Return(models.Url{"testsurl", "testurl"}, nil)

	urlController := New(urlService, config.MustLoad())
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
