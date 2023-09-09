package mocks

import (
	"github.com/en666ki/urlime/internal/shortener/models"
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
	mock "github.com/stretchr/testify/mock"
)

type MockUrlService struct {
	mock.Mock
}

func (m *MockUrlService) StoreShortenUrl(myUrl string) (viewmodels.UrlVM, error) {
	ret := m.Called(myUrl)

	var r0 viewmodels.UrlVM
	if rf, ok := ret.Get(0).(func(string) viewmodels.UrlVM); ok {
		r0 = rf(myUrl)
	} else {
		r0 = ret.Get(0).(viewmodels.UrlVM)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(myUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockUrlService) ReadUrl(surl string) (models.Url, error) {
	ret := m.Called(surl)

	var r0 models.Url
	if rf, ok := ret.Get(0).(func(string) models.Url); ok {
		r0 = rf(surl)
	} else {
		r0 = ret.Get(0).(models.Url)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(surl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
