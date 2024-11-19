package album

import (
	"net/http"
	"piccolo/api/types"
	"piccolo/api/util"

	"github.com/labstack/echo/v4"
)

func (m *AlbumModule) postAlbumPhotosUploadHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := util.GetIdParam(c)
	if albumId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id param.")
	}

	form, err := c.MultipartForm()
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "At least one file is required",
			},
		)
	}

	photoIds, err := m.photoService.UploadFiles(ctx, files, userId)
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	if err = m.albumRepo.InsertPhotos(ctx, albumId, photoIds, userId); err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		types.SuccessRes{
			Success: true,
			Message: "Album photos uploaded",
		},
	)
}
