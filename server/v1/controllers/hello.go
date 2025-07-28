package controllers

import (
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	now := time.Now().Format(time.RFC3339)

	err := rand.IntN(100) < 10 // simulate 10% error
	if err {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed from " + now})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "hello from " + now})
}
