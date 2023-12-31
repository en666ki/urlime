package mocks

import (
	"github.com/en666ki/urlime/internal/shortener/models"
	mock "github.com/stretchr/testify/mock"
)

type MockUrlRepository struct {
	mock.Mock
}

func (m *MockUrlRepository) PutUrl(myUrl, surl string) error {
	ret := m.Called(myUrl, surl)

	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r1 = rf(myUrl, surl)
	} else {
		r1 = ret.Error(0)
	}

	return r1
}

func (m *MockUrlRepository) GetUrl(surl string) (models.Url, error) {
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
