package upload

import (
	"net/http"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (m *UploadModule) postUploadHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.SuccessRes{
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
			Message: "File uploaded",
		},
	)
}
