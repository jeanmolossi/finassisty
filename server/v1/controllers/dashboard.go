package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Dashboard returns a handler that serves the React dashboard index.
func Dashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/dashboard/")
	}
}
