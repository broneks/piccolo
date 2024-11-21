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

func (mod *AuthModule) refreshHandler(c echo.Context) error {
	refreshToken, err := getRefreshTokenString(c)
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	if refreshToken == "" {
		return c.JSON(http.StatusUnauthorized, types.SuccessRes{
			Success: false,
			Message: "Unauthorized: Token is invalid.",
		})
	}

	isValid := jwtoken.VerifyToken(refreshToken)
	if !isValid {
		return c.JSON(http.StatusUnauthorized, types.SuccessRes{
			Success: false,
			Message: "Unauthorized: Token is invalid.",
		})
	}

	userId := jwtoken.GetUserId(refreshToken)
	userEmail := jwtoken.GetUserEmail(refreshToken)

	accessToken, err := jwtoken.NewAccessJwt(userId, userEmail).GenerateToken()
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	setAccessTokenCookie(c, accessToken)

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Token refreshed",
	})
}
