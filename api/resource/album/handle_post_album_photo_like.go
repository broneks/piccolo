package album

import (
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

func (mod *AlbumModule) postAlbumPhotoLike(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	albumId := helper.GetIdParam(c)
	if albumId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid album id param.",
		})
	}

	photoId := helper.GetIdParamByName(c, "photoId")
	if photoId == "" {
		return c.JSON(http.StatusBadRequest, types.SuccessRes{
			Success: false,
			Message: "Invalid photo id param.",
		})
	}

	err := mod.albumRepo.LikePhoto(ctx, albumId, photoId, userId)
	if err != nil {
		switch helper.CheckSqlError(err) {
		case "unique-violation":
			return c.JSON(http.StatusBadRequest, types.SuccessRes{
				Success: false,
				Message: "Already liked",
			})

		default:
			return c.JSON(http.StatusInternalServerError, types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			})
		}
	}

	return nil
}