package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (*sql.DB, error) {
	return New("root", "root", "tasker_dev_db", "tasker_dev")
}

func New(user, pass, host, databaseName string) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		user,
		pass,
		host,
		databaseName,
	)
	return sql.Open("mysql", connStr)
}
