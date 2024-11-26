package api

import (
	"fmt"
	"net/http"
	"piccolo/api/types"
	"strings"

	"github.com/labstack/echo/v4"
)

func httpErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	url := c.Request().URL.String()

	if strings.HasPrefix(url, "/api") && url != "/api/health" {
		c.JSON(code, types.SuccessRes{
			Success: false,
			Message: fmt.Sprintf("%d error", code),
		})
		return
	}

	errorPage := fmt.Sprintf("templates/%d.html", code)

	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}
