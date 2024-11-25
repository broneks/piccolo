package authservice

import (
	"context"
	"fmt"
	"piccolo/api/consts"
	"piccolo/api/service/jwtservice"
)

func (svc *AuthService) UpdateUserPassword(ctx context.Context, token, newPassword string) error {
	if newPassword == "" {
		return fmt.Errorf("New password is missing")
	}

	email := jwtservice.GetUserEmail(token)
	if email == "" {
		return fmt.Errorf("Invalid token")
	}

	user, err := svc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		svc.server.Logger.Error(err.Error())
		return fmt.Errorf("Cannot find user")
	}

	hash, err := svc.hashPassword(newPassword)
	if err != nil {
		svc.server.Logger.Error(err.Error())
		return fmt.Errorf("Cannot hash password")
	}

	err = svc.userRepo.UpdatePassword(ctx, user.Id.String, hash)
	if err != nil {
		svc.server.Logger.Error(err.Error())
		return fmt.Errorf("Cannot update user password")
	}

	err = svc.server.Cache.AddListItems(ctx, consts.ResetPasswordTokenBlacklistKey, token)

	return nil
}
