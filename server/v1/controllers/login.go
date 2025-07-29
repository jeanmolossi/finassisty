package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login renders the login page.
func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]any{
		"Title":       "Acessar",
		"Description": "Entre no Finassisty",
	})
}
