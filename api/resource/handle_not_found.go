package resource

import (
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func handleNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, types.SuccessRes{
		Success: false,
		Message: "not found",
	})
}
