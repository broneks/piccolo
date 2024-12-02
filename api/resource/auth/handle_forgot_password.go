package auth

import (
	"log/slog"
	"net/http"
	"piccolo/api/service/jwtservice"
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
		slog.Error("failed to bind forgot password request data", "err", err)
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
		slog.Error("failed to get user by email", "err", err)
		return c.JSON(http.StatusUnprocessableEntity, types.SuccessRes{
			Success: false,
			Message: "Cannot reset password.",
		})
	}

	resetPasswordToken, err := jwtservice.NewResetPasswordJwt(user.Email.String).GenerateToken()
	if err != nil {
		slog.Error("failed to generate reset password token", "err", err)
	}

	err = mod.server.Mailer.SendResetPassword(ctx, user.Email.String, resetPasswordToken)
	if err != nil {
		slog.Error("failed to send reset password email", "err", err)
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
