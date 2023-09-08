package aerospike

import (
	"testing"

	"github.com/en666ki/urlime/internal/db"
	"github.com/stretchr/testify/suite"
)

type AerospikeTestSuite struct {
	suite.Suite
	a *AerospikeDB
}

func TestAerospike(t *testing.T) {
	suite.Run(t, new(AerospikeTestSuite))
}

func (s *AerospikeTestSuite) SetupTest() {
	a := New("aerospike", 3000, "urlime_test", "url")
	s.NotNil(a)
	s.a = a
}

func (s *AerospikeTestSuite) TestPutUrl() {
	testUrl := db.NewUrl("key", "value", s.a)

	err := s.a.PutUrl(testUrl)
	s.NoError(err)
}

func (s *AerospikeTestSuite) TestGetUrl() {
	testUrl := db.NewUrl("KEY", "VALUE", s.a)

	err := s.a.PutUrl(testUrl)
	s.NoError(err)

	thenUrl, err := s.a.GetUrl("KEY")
	s.NoError(err)
	s.Equal(testUrl, thenUrl)
}

func (s *AerospikeTestSuite) TestPutUrlBadClient() {
	ac := New("0.0.0.0", 3000, "bad", "client")
	testUrl := db.NewUrl("key", "value", ac)

	err := ac.PutUrl(testUrl)
	s.Error(err)
}

func (s *AerospikeTestSuite) TestGetUrlBadClient() {
	ac := New("0.0.0.0", 3000, "bad", "client")

	thenUrl, err := ac.GetUrl("KEY")
	s.Error(err)
	s.Nil(thenUrl)
}
