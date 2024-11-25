package authservice

import (
	"context"
	"fmt"
	"piccolo/api/model"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (svc *AuthService) CreateUser(ctx context.Context, username, email, password string) error {
	if password == "" {
		return fmt.Errorf("Password is missing")
	}

	hash, err := svc.hashPassword(password)
	if err != nil {
		svc.server.Logger.Error(err.Error())
		return fmt.Errorf("Cannot hash password")
	}

	err = svc.userRepo.InsertOne(ctx, model.User{
		Username: pgtype.Text{String: username, Valid: true},
		Email:    pgtype.Text{String: email, Valid: true},
		Hash:     pgtype.Text{String: hash, Valid: true},
		HashedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}
