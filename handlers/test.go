package handlers

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockDBInit(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	// mock DB init
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()

	return db, mock
}
