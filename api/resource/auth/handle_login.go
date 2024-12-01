package auth

import (
	"net/http"
	"piccolo/api/service/jwtservice"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (mod *AuthModule) loginHandler(c echo.Context) error {
	ctx := c.Request().Context()
	ip := c.RealIP()
	req := new(LoginReq)

	var err error

	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	user, err := mod.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	if user == nil {
		mod.banHammerService.RecordFailedAttempt(ctx, ip)
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid email or password.",
		})
	}

	if !mod.authService.VerifyPassword(user.Hash.String, req.Password) {
		mod.banHammerService.RecordFailedAttempt(ctx, ip)
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid email or password.",
		})
	}

	if !user.IsActive() {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "User is not active.",
		})
	}

	if err = mod.userRepo.UpdateLastLoginAt(ctx, user.Id.String); err != nil {
		mod.server.Logger.Error(err.Error())
	}

	accessToken, err := jwtservice.NewAccessJwt(user.Id.String, user.Email.String).GenerateToken()
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	refreshToken, err := jwtservice.NewRefreshJwt(user.Id.String, user.Email.String).GenerateToken()
	if err != nil {
		mod.server.Logger.Error(err.Error())
	}

	c.SetCookie(mod.authService.NewAccessTokenCookie(accessToken))
	c.SetCookie(mod.authService.NewRefreshTokenCookie(refreshToken))

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Logged in",
	})
}
