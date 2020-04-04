package handlers

import (
	"github.com/labstack/echo"
)

// Routing does routing.
func Routing(e *echo.Echo) {

	e.GET("/", index)
}
