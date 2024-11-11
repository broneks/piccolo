package albums

import (
	"net/http"
	"piccolo/api/shared"
	"piccolo/api/util"
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

func (m *AlbumsModule) getAlbumPhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()

	albumId := util.GetIdParam(c)
	if albumId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	photos, _ := m.albumRepo.GetPhotos(ctx, albumId)

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
