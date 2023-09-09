package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/en666ki/urlime/internal/shortener/interfaces/mocks"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/utils"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

func TestStoreShortenUrl(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.On("PutUrl", utils.Shorten("testurl"), "testurl").Return(nil)

	urlService := UrlService{urlRepository}

	expectedUrl := viewmodels.UrlVM{utils.Shorten("testurl"), "testurl"}

	result, err := urlService.StoreShortenUrl("testurl")
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, result)
}

func TestReadUrl(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.On("GetUrl", utils.Shorten("testurl")).Return(models.Url{123, utils.Shorten("testurl"), "testurl"}, nil)

	urlService := UrlService{urlRepository}

	expectedUrl := models.Url{123, utils.Shorten("testurl"), "testurl"}

	result, err := urlService.GetUrl(utils.Shorten("testurl"))
	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, result)
}
