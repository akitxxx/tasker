package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)

	assert.Equal(t, http.StatusOK, rec.Code)
}
