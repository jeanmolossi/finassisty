package main

import (
	"context"
	"finassisty/server/v1/controllers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	api := e.Group("/api/v1")
	api.GET("/hello", controllers.Hello)
	api.GET("/healthcheck", controllers.Healthcheck)

	staticDir := "app/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		staticDir = "app"
	}
	e.Static("/", staticDir)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
