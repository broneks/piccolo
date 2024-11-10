package auth

import (
	"fmt"
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

func (m *AuthModule) refreshHandler(c echo.Context) error {
	refreshToken := c.Request().Header.Get("x-refresh-token")

	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: "Token is missing.",
		})
	}

	isValid := jwtoken.VerifyToken(refreshToken)
	if !isValid {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: "Token is invalid or expired.",
		})
	}

	userId := jwtoken.GetUserId(refreshToken)
	userEmail := jwtoken.GetUserEmail(refreshToken)

	accessToken, err := jwtoken.NewAccessJwt(userId, userEmail).GenerateToken()
	if err != nil {
		m.server.Logger.Error("failed to create jwt access token", err.Error())
		return c.JSON(http.StatusInternalServerError, shared.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	c.Response().Header().Set("authorization", fmt.Sprintf("Bearer %s", accessToken))

	return c.JSON(http.StatusOK, shared.SuccessRes{
		Success: true,
		Message: "Token refreshed",
	})
}
