package upload

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUpload(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	err = uploadFile(file)
	if err != nil {
		return err
	}

	return c.HTML(
		http.StatusOK,
		fmt.Sprintf(
			"<p>File %s uploaded successfully with fields name=%s and email=%s.</p>",
			file.Filename,
			name,
			email,
		),
	)
}

func Router(e *echo.Echo) {
	e.POST("/upload", GetUpload)
}
