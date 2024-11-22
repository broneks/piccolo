package auth

import (
	"net/http"
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

	err = mod.authService.UpdateUserPassword(ctx, req.Token, req.NewPassword)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, types.SuccessRes{
			Success: false,
			Message: "Cannot reset password",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Your password has been reset",
	})
}
