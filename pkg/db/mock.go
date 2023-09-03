package db

import "github.com/stretchr/testify/mock"

type MockDB struct {
	mock.Mock
}

func NewMock() *MockDB {
	return &MockDB{}
}

func (m *MockDB) PutUrl(url *Url) error {
	args := m.Called(url)

	return args.Error(0)
}

func (m *MockDB) GetUrl(surl string) (*Url, error) {
	args := m.Called(surl)

	return args.Get(0).(*Url), args.Error(1)
}
