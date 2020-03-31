package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const port = "5001"

type AuthInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	initRouting(e)

	e.Logger.Fatal(e.Start(":" + port))
}

func initRouting(e *echo.Echo) {
	e.GET("/", index)
	e.POST("/auth", auth)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "hello, tasker!")
}

func auth(c echo.Context) error {
	authInfo := new(AuthInfo)
	if err := c.Bind(authInfo); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authInfo)
}
