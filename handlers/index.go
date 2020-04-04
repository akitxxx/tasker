package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index returns hello message.
func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		str := "hello, tasker"
		return c.JSON(http.StatusOK, str)
	}
}
