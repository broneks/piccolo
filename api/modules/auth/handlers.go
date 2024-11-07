package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *AuthModule) registerHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "TODO")
}
