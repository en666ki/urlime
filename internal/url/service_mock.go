package url

import (
	mock "github.com/stretchr/testify/mock"
)

type MockUrlService struct {
	mock.Mock
}

func (m *MockUrlService) StoreShortenUrl(myUrl string) (UrlVM, error) {
	ret := m.Called(myUrl)

	var r0 UrlVM
	if rf, ok := ret.Get(0).(func(string) UrlVM); ok {
		r0 = rf(myUrl)
	} else {
		r0 = ret.Get(0).(UrlVM)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(myUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockUrlService) ReadUrl(surl string) (Url, error) {
	ret := m.Called(surl)

	var r0 Url
	if rf, ok := ret.Get(0).(func(string) Url); ok {
		r0 = rf(surl)
	} else {
		r0 = ret.Get(0).(Url)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(surl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
