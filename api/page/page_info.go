package page

import (
	"github.com/labstack/echo/v4"
)

type PageInfo struct {
	Nonce string
	Title string
}

func NewPageInfo(c echo.Context, title string) PageInfo {
	nonce := c.Get("nonce").(string)

	return PageInfo{
		Nonce: nonce,
		Title: title,
	}
}
