package albums

import (
	"net/http"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CreateAlbumReq struct {
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description"`
	CoverPhotoId       string `json:"coverPhotoId,omitempty"`
	IsShareLinkEnabled bool   `json:"isShareLinkEnabled"`
}

func (m *AlbumsModule) postAlbumsCreateHandler(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	req := new(CreateAlbumReq)

	var err error

	if err = c.Bind(req); err != nil {
		m.server.Logger.Error(err.Error())
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

	// TODO: Add nullable CoverPhotoId
	err = m.albumRepo.InsertOne(ctx, model.Album{
		UserId:             pgtype.Text{String: userId, Valid: true},
		Name:               pgtype.Text{String: req.Name, Valid: true},
		Description:        pgtype.Text{String: req.Description, Valid: true},
		CoverPhotoId:       pgtype.Text{},
		IsShareLinkEnabled: pgtype.Bool{Bool: req.IsShareLinkEnabled, Valid: true},
	})
	if err != nil {
		m.server.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, types.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	return c.JSON(http.StatusCreated, types.SuccessRes{
		Success: true,
		Message: "Album created",
	})
}
