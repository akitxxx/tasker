package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	str := "hello, tasker"
	return c.JSON(http.StatusOK, str)
}
