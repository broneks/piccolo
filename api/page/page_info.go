package page

import (
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type PageInfo struct {
	Nonce     string
	CsrfToken string
	Title     string
}

func NewPageInfo(c echo.Context, title string) PageInfo {
	env := os.Getenv("ENV")

	nonce := c.Get("nonce").(string)

	var csrfToken string

	if env != "local" {
		csrfToken = c.Get(echoMiddleware.DefaultCSRFConfig.ContextKey).(string)
	}

	return PageInfo{
		Nonce:     nonce,
		CsrfToken: csrfToken,
		Title:     title,
	}
}
