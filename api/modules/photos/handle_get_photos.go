package photos

import (
	"net/http"
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

func (m *PhotosModule) getPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()

	photos, _ := m.photoRepo.GetAll(ctx)

	if len(photos) == 0 {
		return c.JSON(http.StatusOK, []interface{}{})
	}

	var photoResList []PhotoRes

	for _, photo := range photos {
		url := photo.GetUrl(ctx, m.server)

		photoResList = append(photoResList, PhotoRes{
			Id:          photo.Id,
			UserId:      photo.UserId,
			Filename:    photo.Filename,
			FileSize:    photo.FileSize,
			Url:         url,
			ContentType: photo.ContentType,
			CreatedAt:   photo.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, photoResList)
}
