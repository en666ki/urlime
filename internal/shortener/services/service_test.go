package services

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/interfaces/mocks"
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/repositories"
	"github.com/en666ki/urlime/internal/shortener/utils"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

var log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

func TestStoreShortenUrl(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.
		On("PutUrl", utils.Shorten("testurl", utils.Md5Shortener, 8), "testurl").
		Return(nil)

	urlService := New(urlRepository, config.MustLoad(), log)

	expectedUrl := viewmodels.UrlVM{utils.Shorten("testurl", utils.Md5Shortener, 8), "testurl"}

	result := urlService.StoreShortenUrl("testurl")
	vurl := viewmodels.FromModel(models.Url(*result.Data))

	assert.Empty(t, result.Message)
	assert.Equal(t, expectedUrl, vurl)
}

func TestStoreShortenUrlError(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.
		On("PutUrl", utils.Shorten("testurl", utils.Md5Shortener, 8), "testurl").
		Return(repositories.ErrorNotFound)

	urlService := New(urlRepository, config.MustLoad(), log)

	result := urlService.StoreShortenUrl("testurl")

	assert.NotEmpty(t, result.Message)
}

func TestReadUrl(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.
		On("GetUrl", utils.Shorten("testurl", utils.Md5Shortener, 8)).
		Return(models.Url{utils.Shorten("testurl", utils.Md5Shortener, 8), "testurl"}, nil)

	urlService := New(urlRepository, config.MustLoad(), log)

	expectedUrl := viewmodels.UrlVM{utils.Shorten("testurl", utils.Md5Shortener, 8), "testurl"}

	result := urlService.ReadUrl(utils.Shorten("testurl", utils.Md5Shortener, 8))
	vurl := viewmodels.FromModel(models.Url(*result.Data))

	assert.Empty(t, result.Message)
	assert.Equal(t, expectedUrl, vurl)
}

func TestReadUrlError(t *testing.T) {
	urlRepository := new(mocks.MockUrlRepository)

	urlRepository.
		On("GetUrl", utils.Shorten("testurl", utils.Md5Shortener, 8)).
		Return(models.Url{utils.Shorten("testurl", utils.Md5Shortener, 8), "testurl"}, repositories.ErrorKeyExists)

	urlService := New(urlRepository, config.MustLoad(), log)

	result := urlService.ReadUrl(utils.Shorten("testurl", utils.Md5Shortener, 8))

	assert.NotEmpty(t, result.Message)
}
