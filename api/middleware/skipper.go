package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Skipper(c echo.Context) bool {
	env := os.Getenv("ENV")

	return env == "local"
}
