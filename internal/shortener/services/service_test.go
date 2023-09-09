package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreShortenUrl(t *testing.T) {
	urlRepository := new(MockUrlRepository)

	urlRepository.On("PutUrl", shorten("testurl"), "testurl").Return(nil)

	urlService := UrlService{urlRepository}

	expectedUrl := UrlVM{shorten("testurl"), "testurl"}

	result, err := urlService.StoreShortenUrl("testurl")
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, result)
}

func TestReadUrl(t *testing.T) {
	urlRepository := new(MockUrlRepository)

	urlRepository.On("GetUrl", shorten("testurl")).Return(Url{123, shorten("testurl"), "testurl"}, nil)

	urlService := UrlService{urlRepository}

	expectedUrl := Url{123, shorten("testurl"), "testurl"}

	result, err := urlService.GetUrl(shorten("testurl"))
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, result)
}
