package upload

import (
	"net/http"
	"piccolo/api/shared"
	"time"

	"github.com/labstack/echo/v4"
)

type PhotoRes struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	Filename    string    `json:"filename"`
	FileSize    int       `json:"fileSize"`
	Url         string    `json:"url"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (m *UploadModule) getUploadsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photos, _ := m.photoRepo.GetAll(ctx, userId)

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, shared.EmptySlice{})
	}

	var photoResList []PhotoRes

	for _, photo := range photos {
		url := photo.GetUrl(ctx, m.server)

		photoResList = append(photoResList, PhotoRes{
			Id:          photo.Id.String,
			UserId:      photo.UserId.String,
			Filename:    photo.Filename.String,
			FileSize:    int(photo.FileSize.Int32),
			Url:         url,
			ContentType: photo.ContentType.String,
			CreatedAt:   photo.CreatedAt.Time,
		})
	}

	return c.JSON(http.StatusOK, photoResList)
}

func (m *UploadModule) postUploadHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			shared.SuccessRes{
				Success: false,
				Message: "File is required",
			},
		)
	}

	err = m.UploadFile(ctx, file, userId)
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(
			http.StatusBadRequest,
			shared.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		shared.SuccessRes{
			Success: true,
			Message: "File uploaded",
		},
	)
}
