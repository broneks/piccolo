package photo

import (
	"context"
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/model"
	"piccolo/api/types"
	"time"

	"github.com/labstack/echo/v4"
)

func deletePhotoFromObjectStorage(ctx context.Context, mod *PhotoModule, photo *model.Photo, userId string) {
	if photo == nil {
		slog.Error("Error deleting photo in object storage", "err", "photo filename not found")
	}

	err := mod.server.ObjectStorage.DeleteFile(ctx, photo.Filename.String, userId)
	if err != nil {
		slog.Error("Error deleting photo in object storage", "err", err)
	}
}

func (mod *PhotoModule) deletePhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	photoId := helper.GetIdParam(c)
	if photoId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid id param.",
		})
	}

	photo, err := mod.photoRepo.GetById(ctx, photoId, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Photo not found",
		})
	}

	rowsAffected, err := mod.photoRepo.RemoveOne(ctx, photoId, userId)
	if err != nil || rowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "An unexpected error occurred",
		})
	}

	deleteFromObjectStorageCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	go deletePhotoFromObjectStorage(deleteFromObjectStorageCtx, mod, photo, userId)

	return c.JSON(http.StatusOK, types.SuccessRes{
		Success: true,
		Message: "Deleted photo",
	})
}
