package photo

import (
	"log/slog"
	"net/http"
	"piccolo/api/types"
	"strings"

	"github.com/labstack/echo/v4"
)

func (mod *PhotoModule) postPhotosUploadHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	form, err := c.MultipartForm()
	if err != nil {
		slog.Error("failed to grab mulipart form data", "err", err)
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

	if _, err = mod.photoService.UploadFiles(ctx, files, userId); err != nil {
		message := "Unexpected error"

		if strings.Contains(err.Error(), "file storage limit exceeded") {
			message = err.Error()
		}

		slog.Error("failed to upload files to cloud storage", "err", err)
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
				Success: false,
				Message: message,
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		types.SuccessRes{
			Success: true,
			Message: "Photos uploaded",
		},
	)
}
