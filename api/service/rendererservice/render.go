package rendererservice

import (
	"io"

	"github.com/labstack/echo/v4"
)

func (svc *RendererService) Render(w io.Writer, name string, data any, c echo.Context) error {
	if viewContext, isMap := data.(map[string]any); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return svc.templates.ExecuteTemplate(w, name, data)
}
