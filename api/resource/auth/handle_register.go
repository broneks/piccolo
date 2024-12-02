package auth

import (
	"log/slog"
	"net/http"
	"piccolo/api/helper"
	"piccolo/api/types"

	"github.com/labstack/echo/v4"
)

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=14"`
}

func (mod *AuthModule) registerHandler(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(RegisterReq)

	var err error

	if err = c.Bind(req); err != nil {
		slog.Error("failed to bind register user request data", "err", err)
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

	mod.authService.CreateUser(ctx, req.Email, req.Email, req.Password)
	if err != nil {
		switch helper.CheckSqlError(err) {
		case "unique-violation":
			return c.JSON(http.StatusBadRequest, types.SuccessRes{
				Success: false,
				Message: "Email is taken",
			})

		default:
			slog.Error("failed to create user", "err", err)
			return c.JSON(http.StatusInternalServerError, types.SuccessRes{
				Success: false,
				Message: "Unexpected error",
			})
		}
	}

	return c.JSON(http.StatusCreated, types.SuccessRes{
		Success: true,
		Message: "User created",
	})
}
