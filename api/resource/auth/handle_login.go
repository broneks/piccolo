package auth

import (
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (mod *AuthModule) loginHandler(c echo.Context) error {
	ctx := c.Request().Context()
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
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid email or password.",
		})
	}

	if !user.CheckPassword(req.Password) {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid email or password.",
		})
	}

	if err = mod.userRepo.UpdateLastLoginAt(ctx, user.Id.String); err != nil {
		mod.server.Logger.Error(err.Error())
	}

	accessToken, err := jwtoken.NewAccessJwt(user.Id.String, user.Email.String).GenerateToken()
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	refreshToken, err := jwtoken.NewRefreshJwt(user.Id.String, user.Email.String).GenerateToken()
	if err != nil {
		mod.server.Logger.Error(err.Error())
	}

	setAccessTokenCookie(c, accessToken)
	setRefreshTokenCookie(c, refreshToken)

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Logged in",
	})
}
