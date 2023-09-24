package routers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"

	"github.com/en666ki/urlime/internal/config"
	"github.com/en666ki/urlime/internal/shortener/result"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type ShortenerTestSuite struct {
	suite.Suite
	router *chi.Mux
	log    *slog.Logger
}

func TestShortenerTestSuite(t *testing.T) {
	suite.Run(t, new(ShortenerTestSuite))
}

func (s *ShortenerTestSuite) SetupTest() {
	cfg := config.MustLoad()
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	r, err := ShortenRouter().InitRouter(cfg, log)
	s.NoError(err)
	s.router = r
}

func (s *ShortenerTestSuite) TestE2E() {
	randomString := uniuri.New()
	req := httptest.NewRequest("GET", "http://localhost:8080/shorten/"+randomString, nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	expectedData := viewmodels.UrlVM{}
	expectedData.Url = randomString

	actualResult := result.Result{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	s.Equal(expectedData.Url, actualResult.Data.Url)
	s.NotEmpty(actualResult.Data.Surl)
	s.Equal(http.StatusCreated, actualResult.Code)

	surl := actualResult.Data.Surl

	req = httptest.NewRequest("GET", "http://localhost:8080/unshort/"+surl, nil)
	w = httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	expectedData = viewmodels.UrlVM{}
	expectedData.Url = randomString
	expectedData.Surl = surl
	expectedResult := result.Result{
		Data:    &expectedData,
		Code:    http.StatusOK,
		Message: "",
	}

	actualResult = result.Result{}

	json.NewDecoder(w.Body).Decode(&actualResult)
	s.Equal(expectedResult, actualResult)
}
