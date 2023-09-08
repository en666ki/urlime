package url

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	urlService := new(MockUrlService)

	urlService.On("StoreShortenUrl", "testurl").Return(UrlVM{"testsurl", "testurl"}, nil)

	urlController := UrlController{urlService}
	req := httptest.NewRequest("GET", "http://localhost:8080/shorten/testurl", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()

	r.HandleFunc("/shorten/{url}", urlController.Shorten)

	r.ServeHTTP(w, req)

	expecterResult := UrlVM{}
	expecterResult.Url = "testurl"
	expecterResult.Surl = "testsurl"

	actualResult := UrlVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expecterResult, actualResult)
}

func TestUnshort(t *testing.T) {
	urlService := new(MockUrlService)

	urlService.On("ReadUrl", "testsurl").Return(Url{123, "testsurl", "testurl"}, nil)

	urlController := UrlController{urlService}
	req := httptest.NewRequest("GET", "http://localhost:8080/unshort/testsurl", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()

	r.HandleFunc("/unshort/{surl}", urlController.Unshort)

	r.ServeHTTP(w, req)

	expecterResult := UrlVM{}
	expecterResult.Url = "testurl"
	expecterResult.Surl = "testsurl"

	actualResult := UrlVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expecterResult, actualResult)
}
