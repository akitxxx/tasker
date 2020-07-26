package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lelouch99v/tasker/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUserList(t *testing.T) {
	t.Skip()

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

	assert.Equal(t, writer.Code, http.StatusOK)
	assert.Equal(t, users[0].ID, uint64(1))
	assert.Equal(t, users[1].ID, uint64(2))
	assert.Equal(t, users[0].Email, "test1@example.com")
	assert.Equal(t, users[1].Email, "test2@example.com")
}

func Test(t *testing.T) {
	t.Skip()

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

	assert.Equal(t, writer.Code, http.StatusOK)
	assert.Equal(t, user.ID, uint64(1))
	assert.Equal(t, user.Email, "test1@example.com")
}
