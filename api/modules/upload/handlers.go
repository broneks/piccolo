package upload

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *UploadModule) handleGetUploads(c echo.Context) error {
	photos, _ := m.photoRepo.GetAll(c.Request().Context())

	if len(photos) == 0 {
		return c.String(http.StatusOK, "No photos")
	}

	url := photos[0].GetUrl(c.Request().Context(), m.server)

	return c.HTML(
		http.StatusOK,
		fmt.Sprintf(
			"<img src='%s' alt='' />",
			url,
		),
	)
}

func (m *UploadModule) handlePostUpload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	err = m.UploadFile(c.Request().Context(), file)
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
