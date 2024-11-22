package page

import (
	"net/http"
	"piccolo/api/service"

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

func handlePostResetPasswordPage(authService *service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		newPassword := c.FormValue("new-password")
		confirmPassword := c.FormValue("confirm-password")

		error := "Passwords do not match. Please try again."

		// TODO validate incoming passwords

		// TODO call reset password

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
