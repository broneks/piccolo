package upload

import (
	"context"
	"fmt"
	"net/http"
	"piccolo/api/storage/pg"

	"github.com/labstack/echo/v4"
)

func GetUpload(c echo.Context) error {
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
			"<p>File %s uploaded successfully.</p>",
			file.Filename,
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

		url := photos[0].GetUrl(c)

		return c.HTML(
			http.StatusOK,
			fmt.Sprintf(
				"<img src='%s' alt='' />",
				url,
			),
		)
	})

	e.POST("/upload", GetUpload)
}
