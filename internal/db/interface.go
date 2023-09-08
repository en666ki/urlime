package db

import "database/sql"

type IDHandler interface {
	Execute(statement string) (IResult, error)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}

type IResult sql.Result
