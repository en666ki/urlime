package infrastructures

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type PostgresqlHandlerSuite struct {
	suite.Suite
	handler PostgresqlHandler
}

func TestPostgresqlHandler(t *testing.T) {
	suite.Run(t, new(PostgresqlHandlerSuite))
}

func (s *PostgresqlHandlerSuite) SetupTest() {
	sqlConn, err := sql.Open("postgres", "host=postgres_test port=5432 dbname=local user=local password=local_pwd sslmode=disable")
	s.NoError(err)
	s.NotNil(sqlConn)
	s.handler.Conn = sqlConn
}

func (s *PostgresqlHandlerSuite) TearDownTest() {
	_, err := s.handler.Execute(`
		TRUNCATE TABLE local_urls;
	`)
	s.NoError(err)
	err = s.handler.Conn.Close()
	s.NoError(err)
}

func (s *PostgresqlHandlerSuite) TestExecute() {
	result, err := s.handler.Execute(`
		CALL test_proc();
	`)
	s.NoError(err)
	ra, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(int64(0), ra)
}

func (s *PostgresqlHandlerSuite) TestQuery() {
	result, err := s.handler.Execute(`
		INSERT INTO local_urls (surl, url) VALUES ('tst', 'test');
	`)
	s.NoError(err)
	ra, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(int64(1), ra)
	row, err := s.handler.Query(`
		SELECT * FROM local_urls;
	`)
	s.NoError(err)
	var surl, url string
	r := row.Next()
	s.True(r)
	err = row.Scan(new(int64), &surl, &url)
	s.NoError(err)
	s.Equal("tst", surl)
	s.Equal("test", url)
}
