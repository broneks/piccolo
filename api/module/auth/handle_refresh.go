package auth

import (
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func getRefreshTokenString(c echo.Context) (string, error) {
	var err error

	// first try getting the jwt via an http-only cookie
	tokenCookie, err := c.Cookie("piccolo-refresh-token")
	if err != nil {
		return "", err
	}

	if tokenCookie != nil {
		return tokenCookie.Value, nil
	}

	// fallback to using the auth header
	tokenString, err := jwtoken.ExtractTokenString(c.Request().Header.Get("x-refresh-token"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (m *AuthModule) refreshHandler(c echo.Context) error {
	refreshToken, err := getRefreshTokenString(c)
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Token is missing.",
		})
	}

	isValid := jwtoken.VerifyToken(refreshToken)
	if !isValid {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Token is invalid or expired.",
		})
	}

	userId := jwtoken.GetUserId(refreshToken)
	userEmail := jwtoken.GetUserEmail(refreshToken)

	accessToken, err := jwtoken.NewAccessJwt(userId, userEmail).GenerateToken()
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	// TODO is this needed?
	// c.Response().Header().Set("authorization", fmt.Sprintf("Bearer %s", accessToken))

	setAccessTokenCookie(c, accessToken)

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Token refreshed",
	})
}
