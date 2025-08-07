package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, ping())
	})

	e.GET("/easy", func(c echo.Context) error {
		return c.String(http.StatusOK, easy())
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func ping() string {
	return "pong"
}

func easy() string {
	return "easy!!!"
}
