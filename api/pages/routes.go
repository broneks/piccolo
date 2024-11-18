package pages

import (
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, server *types.Server) {
	e.GET("/test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test.html", map[string]string{
			"name": "World!",
		})
	})
}
