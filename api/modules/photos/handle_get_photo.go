package photos

import (
	"net/http"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *PhotosModule) getPhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := util.GetIdParam(c)
	if photoId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	photo, err := m.photoRepo.GetById(ctx, photoId, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	url := photo.GetUrl(ctx, m.server)

	photoRes := PhotoRes{
		Id:          photo.Id.String,
		UserId:      photo.UserId.String,
		Filename:    photo.Filename.String,
		FileSize:    int(photo.FileSize.Int32),
		Url:         url,
		ContentType: photo.ContentType.String,
		CreatedAt:   photo.CreatedAt.Time,
	}

	return c.JSON(http.StatusOK, photoRes)
}
