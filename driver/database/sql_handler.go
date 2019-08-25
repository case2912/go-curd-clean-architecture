package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/case2912/go-curd-clean-architecture/interface/adapter"
)

type SQLHandler struct {
	Conn *sql.DB
}

func NewSQLHandler(config *SQLConfig) adapter.SQLHandler {
	conn, err := sql.Open(config.SQLName, config.GetStr())

	if err != nil {
		err = fmt.Errorf("Initialize database connection is failed:\n%s", err.Error())
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		err = fmt.Errorf("Invalid database connection:\n%s", err.Error())
		log.Fatal(err)
	}
	handler := new(SQLHandler)
	handler.Conn = conn
	return handler
}

func (handler *SQLHandler) Query(statement string, args ...interface{}) (adapter.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}
