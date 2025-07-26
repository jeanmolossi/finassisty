package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	slog.Info("http request")
	return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
}
