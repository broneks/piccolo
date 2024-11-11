package auth

import (
	"net/http"
	"piccolo/api/model"
	"piccolo/api/shared"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=14"`
}

func (m *AuthModule) registerHandler(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(RegisterReq)

	var err error

	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: "Invalid input",
		})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: err.Error(),
		})
	}

	hash, err := hashPassword(req.Password)
	if err != nil {
		m.server.Logger.Error("unexpected error", err)
		return c.JSON(http.StatusInternalServerError, shared.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	err = m.userRepo.InsertOne(ctx, model.User{
		Username: pgtype.Text{String: req.Email},
		Email:    pgtype.Text{String: req.Email},
		Hash:     pgtype.Text{String: hash},
		HashedAt: pgtype.Timestamptz{Time: time.Now()},
	})
	if err != nil {
		m.server.Logger.Error("unexpected error", err)
		return c.JSON(http.StatusInternalServerError, shared.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	return c.JSON(http.StatusOK, shared.SuccessRes{
		Success: true,
		Message: "User created",
	})
}
