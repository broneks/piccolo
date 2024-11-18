package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	errorPage := fmt.Sprintf("templates/%d.html", code)

	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}
