package upload

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type PhotoRes struct {
	Id          string    `json:"id"`
	Filename    string    `json:"filename"`
	FileSize    int       `json:"fileSize"`
	Url         string    `json:"url"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (m *UploadModule) getUploadsHandler(c echo.Context) error {
	photos, _ := m.photoRepo.GetAll(c.Request().Context())

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []interface{}{})
	}

	var photoResList []PhotoRes

	for _, photo := range photos {
		url := photo.GetUrl(c.Request().Context(), m.server)

		photoResList = append(photoResList, PhotoRes{
			Id:          photo.Id,
			Filename:    photo.Filename,
			FileSize:    photo.FileSize,
			Url:         url,
			ContentType: photo.ContentType,
			CreatedAt:   photo.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, photoResList)
}

func (m *UploadModule) handleGetUploads(c echo.Context) error {
	photos, _ := m.photoRepo.GetAll(c.Request().Context())

	if len(photos) == 0 {
		return c.String(http.StatusOK, "No photos")
	}

	var imageTags []string

	for _, photo := range photos {
		url := photo.GetUrl(c.Request().Context(), m.server)

		imageTags = append(imageTags, fmt.Sprintf(
			"<img src='%s' alt='' />",
			url,
		))
	}

	return c.HTML(
		http.StatusOK,
		strings.Join(imageTags, "<br/>"),
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
