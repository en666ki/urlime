package infrastructures

import (
	"database/sql"

	"github.com/en666ki/urlime/internal/db"
)

type PostgresqlHandler struct {
	Conn *sql.DB
}

func (handler *PostgresqlHandler) Execute(statement string) (db.IResult, error) {
	return handler.Conn.Exec(statement)
}

func (handler *PostgresqlHandler) Query(statement string) (db.IRow, error) {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		return new(PostgressRow), err
	}

	row := new(PostgressRow)
	row.Rows = rows

	return row, nil
}

type PostgressRow struct {
	Rows *sql.Rows
}

func (r PostgressRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func (r PostgressRow) Next() bool {
	return r.Rows.Next()
}
