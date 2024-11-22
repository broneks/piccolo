package auth

import (
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

type ResetPasswordReq struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"password" validate:"required,min=14"`
}

func (mod *AuthModule) resetPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(ResetPasswordReq)

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

	email := jwtoken.GetUserEmail(req.Token)
	if email == "" {
		return c.JSON(http.StatusUnprocessableEntity, types.SuccessRes{
			Success: false,
			Message: "Cannot reset password.",
		})
	}

	user, err := mod.userRepo.GetByEmail(ctx, email)
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, types.SuccessRes{
			Success: false,
			Message: "Cannot reset password.",
		})
	}

	hash, err := hashPassword(req.NewPassword)
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	err = mod.userRepo.UpdatePassword(ctx, user.Id.String, hash)
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Your password has been reset",
	})
}
