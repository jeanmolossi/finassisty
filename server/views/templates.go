// Package views contains server-side HTML templates.
package views

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed *.html
var templateFS embed.FS

// TemplateRenderer renders Go HTML templates.
type TemplateRenderer struct {
	templates *template.Template
}

// NewRenderer parses embedded templates and returns a renderer.
func NewRenderer() (*TemplateRenderer, error) {
	tmpl, err := template.ParseFS(templateFS, "*.html")
	if err != nil {
		return nil, err
	}
	return &TemplateRenderer{templates: tmpl}, nil
}

// Render satisfies echo.Renderer.
func (t *TemplateRenderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
