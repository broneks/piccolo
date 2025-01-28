package album

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

type CreateAlbumPhotoReq struct {
	PhotoId string `json:"photoId" validate:"required"`
}

func (mod *AlbumModule) postAlbumPhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	req := new(CreateAlbumPhotoReq)

	var err error

	if err = c.Bind(req); err != nil {
		slog.Error("failed to bind create album photo request data", "err", err)
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid input",
		})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid album id param.",
		})
	}

	canRead, err := mod.photoRepo.CanReadPhoto(ctx, req.PhotoId, userId)
	if err != nil {
		slog.Error("error checking photo read access", "err", err)
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}
	if !canRead {
		return c.JSON(http.StatusNotFound, types.SuccessRes{
			Success: false,
			Message: "Not found",
		})
	}

	if err = mod.albumRepo.InsertPhotos(ctx, albumId, []string{req.PhotoId}, userId); err != nil {
		switch helper.CheckSqlError(err) {
		case "unique-violation":
			return c.JSON(http.StatusBadRequest, types.SuccessRes{
				Success: false,
				Message: "Album photo already exists",
			})

		default:
			slog.Error("error creating album photo", "err", err)
			return c.JSON(http.StatusInternalServerError, types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			})
		}
	}

	return c.JSON(
		http.StatusCreated,
		types.SuccessRes{
			Success: true,
			Message: "Album photo created",
		},
	)
}
