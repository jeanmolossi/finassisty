package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"finassisty/server/config"
	"finassisty/server/infrastructure/support/telemetry"
	"finassisty/server/v1/controllers"
	"finassisty/server/v1/middlewares"
	"finassisty/server/views"
)

func main() {
	config.Env() // load env

	ctx := context.Background()

	err := telemetry.StartCollectors(ctx)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middlewares.OTel())
	e.Use(middlewares.MetrifyRequest())
	e.Use(middlewares.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	renderer, err := views.NewRenderer()
	if err != nil {
		panic(err)
	}
	e.Renderer = renderer

	api := e.Group("/api/v1")
	api.GET("/hello", controllers.Hello)
	api.GET("/healthcheck", controllers.Healthcheck)
	api.GET("/auth/google/login", controllers.GoogleLogin)
	api.GET("/auth/google/callback", controllers.GoogleCallback)
	api.GET("/auth/session", controllers.CheckSession)

	staticDir := "app/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		staticDir = "app"
	}

	e.GET("/", controllers.Home)
	e.GET("/acessar", controllers.Login)
	e.GET("/dashboard", controllers.Dashboard(staticDir), middlewares.RequireSession())

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
