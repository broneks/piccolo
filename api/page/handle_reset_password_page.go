package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResetPasswordPayload struct {
	// TODO
}

func handleResetPasswordPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "reset_password.html", &ResetPasswordPayload{})
	}
}
