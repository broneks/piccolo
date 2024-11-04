package upload

import (
	"context"
	"fmt"
	"net/http"
	"piccolo/api/storage/pg"

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
	e.GET("/uploads", func(c echo.Context) error {
		db := pg.Client(context.Background())
		photos, _ := db.GetPhotos(context.Background())

		if len(photos) == 0 {
			return c.String(http.StatusOK, "No photos")
		}

		url := photos[0].GetUrl()

		return c.HTML(
			http.StatusOK,
			fmt.Sprintf(
				"<img src='%s' alt='' width='1000' />",
				url,
			),
		)
	})

	e.POST("/upload", GetUpload)
}
