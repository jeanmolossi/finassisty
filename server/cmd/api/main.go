package main

import (
	"finassisty/server/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// API route example
	e.GET("/api/hello", controllers.Hello)

	// serve frontend
	staticDir := "app/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		staticDir = "app" // for dev server before build
	}
	e.Static("/", staticDir)

	e.Logger.Fatal(e.Start(":8080"))
}
