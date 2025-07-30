package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CheckSession verifies if a session cookie is present.
func CheckSession(c echo.Context) error {
	sess, err := c.Cookie("session")
	if err != nil || sess.Value == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	return c.NoContent(http.StatusOK)
}
