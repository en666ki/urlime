package aerospike

import (
	"testing"

	"github.com/en666ki/urlime/internal/db"
	"github.com/stretchr/testify/suite"
)

type DBAerospikeTestSuite struct {
	suite.Suite
	db *AerospikeDB
}

func TestDBAerospike(t *testing.T) {
	suite.Run(t, new(DBAerospikeTestSuite))
}

func (s *DBAerospikeTestSuite) SetupTest() {
	s.db = New("aerospike", 3000, "urlime_test", "url")
}

func (s *DBAerospikeTestSuite) TestPut() {
	url := db.NewUrl("z", "v", s.db)
	err := url.Put()
	s.NoError(err)
}

func (s *DBAerospikeTestSuite) TestGet() {
	url := db.NewUrl("q", "r", s.db)
	err := url.Put()
	s.NoError(err)
	oldUrl := url
	readUrl, err := url.Get()
	s.NoError(err)
	s.Equal(oldUrl, url)
	s.Equal(oldUrl, readUrl)
}
