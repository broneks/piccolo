package auth

import (
	"fmt"
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (m *AuthModule) loginHandler(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(LoginReq)

	var err error

	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := m.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, shared.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	if user == nil {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: "Invalid email or password.",
		})
	}

	if !user.CheckPassword(req.Password) {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: "Invalid email or password.",
		})
	}

	if err = m.userRepo.UpdateLastLoginAt(ctx, user.Id); err != nil {
		m.server.Logger.Error("failed to update last login at", err.Error())
	}

	accessToken, err := jwtoken.NewAccessJwt(user.Id, user.Email).GenerateToken()
	if err != nil {
		m.server.Logger.Error("failed to create jwt access token", err.Error())
		return c.JSON(http.StatusInternalServerError, shared.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	refreshToken, err := jwtoken.NewRefreshJwt(user.Id, user.Email).GenerateToken()
	if err != nil {
		m.server.Logger.Error("failed to create jwt refresh token", err.Error())
	}

	c.Response().Header().Set("authorization", fmt.Sprintf("Bearer %s", accessToken))
	c.Response().Header().Set("x-refresh-token", fmt.Sprintf("Bearer %s", refreshToken))

	return c.JSON(http.StatusOK, shared.SuccessRes{
		Success: true,
		Message: "Logged in",
	})
}
