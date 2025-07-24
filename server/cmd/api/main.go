package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API route example
	e.GET("/api/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
	})

	// serve frontend
	staticDir := "app/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		staticDir = "app" // for dev server before build
	}
	e.Static("/", staticDir)

	e.Logger.Fatal(e.Start(":8080"))
}
