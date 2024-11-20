package auth

import (
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

type ForgotPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}

func (mod *AuthModule) forgotPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(ForgotPasswordReq)

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
			Message: "Invalid email.",
		})
	}

	resetPasswordToken, err := jwtoken.NewResetPasswordJwt(user.Email.String).GenerateToken()
	if err != nil {
		mod.server.Logger.Error(err.Error())
	}

	baseUrl := c.Request().URL.String()

	err = mod.server.Mailer.SendResetPassword(ctx, user.Email.String, baseUrl, resetPasswordToken)
	if err != nil {
		mod.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "A reset password link has been sent to your email.",
	})
}
