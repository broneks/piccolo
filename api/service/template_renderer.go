package service

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (tr *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	if viewContext, isMap := data.(map[string]any); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return tr.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(dirPath string) *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(dirPath)),
	}
}
