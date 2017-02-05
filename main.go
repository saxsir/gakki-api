package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Image struct {
	URL string
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []*Image{
			&Image{URL: "aaa"},
			&Image{URL: "bbb"},
		})
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
