package photos

import (
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (m *PhotosModule) postPhotosUploadHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

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

	err = m.UploadFiles(ctx, files, userId)
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

	return c.JSON(
		http.StatusOK,
		types.SuccessRes{
			Success: true,
			Message: "Files uploaded",
		},
	)
}
