package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DBTestSuite struct {
	suite.Suite
}

func TestDB(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}

func (s *DBTestSuite) TestNew() {
	db := NewMock()
	defer func() { db.AssertExpectations(s.T()) }()
	url := NewUrl("key", "value", db)
	s.Equal(&Url{"key", "value", db}, url)
}

func (s *DBTestSuite) TestPut() {
	db := NewMock()
	defer func() { db.AssertExpectations(s.T()) }()
	url := NewUrl("z", "v", db)
	db.On("PutUrl", url).Return(nil).Once()
	err := url.Put()
	s.NoError(err)
}

func (s *DBTestSuite) TestGet() {
	db := NewMock()
	defer func() { db.AssertExpectations(s.T()) }()
	url := NewUrl("q", "r", db)
	db.On("GetUrl", "q").Return(url, nil).Once()
	oldUrl := url
	readUrl, err := url.Get()
	s.NoError(err)
	s.Equal(oldUrl, url)
	s.Equal(oldUrl, readUrl)
}
