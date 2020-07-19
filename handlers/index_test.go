package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {

	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, rec.Code)
}
