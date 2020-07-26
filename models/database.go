package models

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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
