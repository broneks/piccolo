package modules

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "not found",
	})
}
