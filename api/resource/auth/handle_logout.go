package auth

import (
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AuthModule) logoutHandler(c echo.Context) error {
	accessCookie, err := c.Cookie("piccolo-access-token")
	if err != nil || accessCookie == nil {
		mod.server.Logger.Debug(err.Error())
		return c.JSON(http.StatusForbidden, types.SuccessRes{
			Success: false,
			Message: "Forbidden",
		})
	}
	if accessCookie.Value == "" {
		return c.JSON(http.StatusForbidden, types.SuccessRes{
			Success: false,
			Message: "Forbidden",
		})
	}

	c.SetCookie(mod.authService.NewAccessTokenCookie(""))
	c.SetCookie(mod.authService.NewRefreshTokenCookie(""))

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Successfully logged out.",
	})
}
