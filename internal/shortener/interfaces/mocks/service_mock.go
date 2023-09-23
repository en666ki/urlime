package mocks

import (
	"github.com/en666ki/urlime/internal/shortener/result"
	mock "github.com/stretchr/testify/mock"
)

type MockUrlService struct {
	mock.Mock
}

func (m *MockUrlService) StoreShortenUrl(myUrl string) result.Result {
	ret := m.Called(myUrl)

	var r0 result.Result
	if rf, ok := ret.Get(0).(func(string) result.Result); ok {
		r0 = rf(myUrl)
	} else {
		r0 = ret.Get(0).(result.Result)
	}
	return r0
}

func (m *MockUrlService) ReadUrl(surl string) result.Result {
	ret := m.Called(surl)

	var r0 result.Result
	if rf, ok := ret.Get(0).(func(string) result.Result); ok {
		r0 = rf(surl)
	} else {
		r0 = ret.Get(0).(result.Result)
	}
	return r0
}
