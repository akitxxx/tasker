package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestHandleAuth(t *testing.T) {
	t.Skip()
	_, mock := MockDBInit(t)

	// define mock
	email := "test@test.com"
	password := "password"
	columns := []string{"id", "email"}
	mock.ExpectQuery("select id, email from users where email = \\? and password = \\?").
		WithArgs(email, password).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, email))

	// execute
	mux := http.NewServeMux()           // テストを実行するマルチプレクサを生成
	mux.HandleFunc("/auth", HandleAuth) // テスト対象のハンドラを付加

	jsonStr := strings.NewReader(fmt.Sprintf(`{"email":"%s", "password":"%s"}`, email, password))

	writer := httptest.NewRecorder()                        // 返されたhttp レスポンスを取得
	request, _ := http.NewRequest("POST", "/auth", jsonStr) // テストしたいハンドラ宛のリクエストを作成
	mux.ServeHTTP(writer, request)                          // テスト対象のハンドラにリクエストを送信

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet: %s", err)
	}

	assert.Equal(t, 200, writer.Code)
}
