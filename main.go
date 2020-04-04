package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lelouch99v/tasker/handlers"
)

const port = "5001"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routing(e)

	e.Start(":" + port)
}

func routing(e *echo.Echo) {
	e.GET("/", handlers.Index())
}
