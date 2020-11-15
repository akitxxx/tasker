package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/lelouch99v/tasker/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUserList(t *testing.T) {
	t.Skip()
	_, mock := MockDBInit(t)

	// define mock
	columns := []string{"id", "email"}
	email1 := "test1@test.com"
	email2 := "test2@test.com"
	mock.ExpectQuery("select id, email from users;").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, email1).AddRow(2, email2))

	// execute
	mux := http.NewServeMux()
	mux.HandleFunc("/user", GetUserList)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user", nil)
	mux.ServeHTTP(writer, request)

	var users []models.User
	err := json.Unmarshal(writer.Body.Bytes(), &users)
	if err != nil {
		t.Errorf(err.Error())
	}

	// assert
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet: %s", err)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, uint64(1), users[0].ID)
	assert.Equal(t, uint64(2), users[1].ID)
	assert.Equal(t, email1, users[0].Email)
	assert.Equal(t, email2, users[1].Email)
}

func TestFindUserById(t *testing.T) {
	t.Skip()
	_, mock := MockDBInit(t)

	// define mock
	columns := []string{"id", "email"}
	email := "test1@test.com"
	mock.ExpectQuery("select id, email from users where id = ?;").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, email))

	// execute
	mux := http.NewServeMux()
	mux.HandleFunc("/user/", FindUserById)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/1", nil)
	mux.ServeHTTP(writer, request)

	var user models.User
	err := json.Unmarshal(writer.Body.Bytes(), &user)
	if err != nil {
		t.Errorf(err.Error())
	}

	// asssert
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet: %s", err)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, uint64(1), user.ID)
	assert.Equal(t, email, user.Email)
}
