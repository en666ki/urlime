package postgresql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/en666ki/urlime/internal/config"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type PostgresqlHandlerSuite struct {
	suite.Suite
	handler PostgresqlHandler
	cfg     *config.Config
}

func TestPostgresqlHandler(t *testing.T) {
	suite.Run(t, new(PostgresqlHandlerSuite))
}

func (s *PostgresqlHandlerSuite) SetupTest() {
	s.cfg = config.MustLoad()
	sqlConn, err := sql.Open(s.cfg.DB.Driver,
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			s.cfg.DB.Host, s.cfg.DB.Port, s.cfg.DB.Name, s.cfg.DB.User, s.cfg.DB.Password, s.cfg.DB.SslMode))
	s.NoError(err)
	s.NotNil(sqlConn)
	s.handler.Conn = sqlConn
}

func (s *PostgresqlHandlerSuite) TearDownTest() {
	_, err := s.handler.Execute(fmt.Sprintf(`
		TRUNCATE TABLE %s;
	`, s.cfg.DB.Table))
	s.NoError(err)
	err = s.handler.Conn.Close()
	s.NoError(err)
}

func (s *PostgresqlHandlerSuite) TestExecute() {
	result, err := s.handler.Execute(fmt.Sprintf(`
		INSERT INTO %s (surl, url) VALUES ('tst_exec', 'test_exec');
	`, s.cfg.DB.Table))
	s.NoError(err)
	ra, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(int64(1), ra)
}

func (s *PostgresqlHandlerSuite) TestQuery() {
	result, err := s.handler.Execute(fmt.Sprintf(`
		INSERT INTO %s (surl, url) VALUES ('tst', 'test');
	`, s.cfg.DB.Table))
	s.NoError(err)
	ra, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(int64(1), ra)
	row, err := s.handler.Query(fmt.Sprintf(`
		SELECT * FROM %s;
	`, s.cfg.DB.Table))
	s.NoError(err)
	var surl, url string
	r := row.Next()
	s.True(r)
	err = row.Scan(&surl, &url)
	s.NoError(err)
	s.Equal("tst", surl)
	s.Equal("test", url)
}
