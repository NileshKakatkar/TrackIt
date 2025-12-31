package conn

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ExecResult struct {
	InsertId     int64
	RowsAffected int64
	Error        error
}

type QueryResult struct {
	Rows  *sql.Rows
	Error error
}
