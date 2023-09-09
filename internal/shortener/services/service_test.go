package services

import (
	"testing"

	"github.com/en666ki/urlime/internal/shortener/interfaces/mocks"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
	"github.com/stretchr/testify/assert"
)

func TestStoreShortenUrl(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.On("PutUrl", shorten("testurl"), "testurl").Return(nil)

	urlService := UrlService{urlRepository}

	expectedUrl := viewmodels.UrlVM{shorten("testurl"), "testurl"}

	result, err := urlService.StoreShortenUrl("testurl")
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, result)
}

func TestReadUrl(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.On("GetUrl", shorten("testurl")).Return(models.Url{123, shorten("testurl"), "testurl"}, nil)

	urlService := UrlService{urlRepository}

	expectedUrl := models.Url{123, shorten("testurl"), "testurl"}

	result, err := urlService.GetUrl(shorten("testurl"))
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, result)
}
