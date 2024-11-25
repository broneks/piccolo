package page

import (
	"fmt"
	"net/http"
	"piccolo/api/service/authservice"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
)

type ResetPasswordPayload struct {
	PageInfo
	Token           string
	Success         bool
	Error           string
	NewPassword     string
	ConfirmPassword string
}

func handleGetResetPasswordPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")

		pageInfo := NewPageInfo(c, "Reset Password")

		return c.Render(http.StatusOK, "reset_password.html", &ResetPasswordPayload{
			PageInfo:        pageInfo,
			Token:           token,
			Success:         false,
			Error:           "",
			NewPassword:     "",
			ConfirmPassword: "",
		})
	}
}

func handlePostResetPasswordPage(authService *authservice.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token := c.FormValue("token")
		newPassword := c.FormValue("new-password")
		confirmPassword := c.FormValue("confirm-password")

		var error string
		var success bool

		if newPassword == "" {
			error = "New password is required."
		} else if utf8.RuneCountInString(newPassword) < authService.MinPasswordCharLength {
			error = fmt.Sprintf("New password is too short. Must be at least %d characters.", authService.MinPasswordCharLength)
		} else if confirmPassword == "" {
			error = "Confirm password is required."
		} else if newPassword != confirmPassword {
			error = "Passwords do not match."
		} else {
			err := authService.UpdateUserPassword(ctx, token, newPassword)
			if err != nil {
				error = "An unexpected error occurred."
			} else {
				success = true
			}
		}

		pageInfo := NewPageInfo(c, "Reset Password")

		return c.Render(http.StatusOK, "reset_password.html", &ResetPasswordPayload{
			PageInfo:        pageInfo,
			Token:           token,
			Success:         success,
			Error:           error,
			NewPassword:     newPassword,
			ConfirmPassword: confirmPassword,
		})
	}
}
