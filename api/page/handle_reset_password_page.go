package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResetPasswordPayload struct {
	PageInfo
	Token           string
	Error           string
	NewPassword     string
	ConfirmPassword string
}

func handleGetResetPasswordPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")

		return c.Render(http.StatusOK, "reset_password.html", &ResetPasswordPayload{
			PageInfo: PageInfo{
				Title: "Reset Password",
			},
			Token:           token,
			Error:           "",
			NewPassword:     "",
			ConfirmPassword: "",
		})
	}
}

func handlePostResetPasswordPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		newPassword := c.FormValue("new-password")
		confirmPassword := c.FormValue("confirm-password")

		error := "Passwords do not match. Please try again."

		return c.Render(http.StatusOK, "reset_password.html", &ResetPasswordPayload{
			PageInfo: PageInfo{
				Title: "Reset Password",
			},
			Token:           token,
			Error:           error,
			NewPassword:     newPassword,
			ConfirmPassword: confirmPassword,
		})
	}
}
