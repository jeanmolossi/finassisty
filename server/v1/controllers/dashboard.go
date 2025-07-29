package controllers

import (
	"path"

	"github.com/labstack/echo/v4"
)

// Dashboard returns a handler that serves the React dashboard index.
func Dashboard(staticDir string) echo.HandlerFunc {
	index := path.Join(staticDir, "index.html")
	return func(c echo.Context) error {
		return c.File(index)
	}
}
