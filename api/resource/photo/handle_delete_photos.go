package photo

import (
	"context"
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func deletePhotosFromObjectStorage(ctx context.Context, mod *PhotoModule, photos []model.Photo, photoIds []string, userId string) {
	for _, photoId := range photoIds {
		var photo *model.Photo

		for _, p := range photos {
			if p.Id.String == photoId {
				photo = &p
				break
			}
		}

		if photo == nil {
			slog.Error("Error deleting photo in object storage", "err", "photo filename not found")
		}

		err := mod.server.ObjectStorage.DeleteFile(ctx, photo.Filename.String, userId)
		if err != nil {
			slog.Error("Error deleting photo in object storage", "err", err)
		}
	}
}

func (mod *PhotoModule) deletePhotosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoIds := helper.GetListParam(c, "ids")
	if len(photoIds) == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Missing ids param.",
		})
	}

	photos, err := mod.photoRepo.GetByIds(ctx, photoIds, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error occurred",
		})
	}
	if len(photos) == 0 {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Photos not found",
		})
	}

	rowsAffected, err := mod.photoRepo.RemoveMany(ctx, photoIds, userId)
	if err != nil || rowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error occurred",
		})
	}

	go deletePhotosFromObjectStorage(context.Background(), mod, photos, photoIds, userId)

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Deleted photos",
	})
}
