package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleAuth(t *testing.T) {
	mux := http.NewServeMux()           // テストを実行するマルチプレクサを生成
	mux.HandleFunc("/auth", HandleAuth) // テスト対象のハンドラを付加

	jsonStr := strings.NewReader(`{"email":"test@test.com", "password":"test"}`)

	writer := httptest.NewRecorder()                        // 返されたhttp レスポンスを取得
	request, _ := http.NewRequest("POST", "/auth", jsonStr) // テストしたいハンドラ宛のリクエストを作成
	mux.ServeHTTP(writer, request)                          // テスト対象のハンドラにリクエストを送信

	var auth Auth
	json.Unmarshal(writer.Body.Bytes(), &auth)
	assert.Equal(t, writer.Code, 200)
	assert.Equal(t, "test@test.com", auth.Email)
	assert.Equal(t, "test", auth.Password)
}
