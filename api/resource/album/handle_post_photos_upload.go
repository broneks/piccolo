package album

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) postAlbumPhotosUploadHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		slog.Error("failed to grab multipart form data", "err", err)
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

	photoIds, err := mod.photoService.UploadFiles(ctx, files, userId)
	if err != nil {
		slog.Error("failed to upload photos", "err", err)
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			},
		)
	}

	if err = mod.albumRepo.InsertPhotos(ctx, albumId, photoIds, userId); err != nil {
		slog.Error("failed to insert photos", "err", err)
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
