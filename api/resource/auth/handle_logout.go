package auth

import (
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AuthModule) logoutHandler(c echo.Context) error {
	c.SetCookie(mod.authService.NewAccessTokenCookie(""))
	c.SetCookie(mod.authService.NewRefreshTokenCookie(""))

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Successfully logged out.",
	})
}
