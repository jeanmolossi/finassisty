package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home renders the landing page.
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]any{
		"Title":       "Finassisty",
		"Description": "Gerencie suas finanças de forma simples",
	})
}
