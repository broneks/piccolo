package service

import (
	"context"
	"fmt"
	"net/http"
	"piccolo/api/jwtoken"
	"piccolo/api/model"
	"piccolo/api/repo"
	"piccolo/api/types"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	server   *types.Server
	userRepo *repo.UserRepo
}

func NewAuthService(server *types.Server, userRepo *repo.UserRepo) *AuthService {
	return &AuthService{
		server:   server,
		userRepo: userRepo,
	}
}

const COST = bcrypt.DefaultCost + 2

func (svc *AuthService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (svc *AuthService) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func (svc *AuthService) NewAccessTokenCookie(value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "piccolo-access-token",
		Value:    value,
		HttpOnly: true,
		Secure:   false, // TODO change this for production
		// SameSite: http.SameSiteStrictMode, // Prevents CSRF by restricting cross-site cookie transmission TODO
		Path: "/",
	}

	if value == "" {
		cookie.Expires = time.Unix(0, 0)
	}

	return cookie
}

func (svc *AuthService) NewRefreshTokenCookie(value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "piccolo-refresh-token",
		Value:    value,
		HttpOnly: true,
		Secure:   false, // TODO change this for production
		// SameSite: http.SameSiteStrictMode, // Prevents CSRF by restricting cross-site cookie transmission TODO
		Path: "/",
	}

	if value == "" {
		cookie.Expires = time.Unix(0, 0)
	}

	return cookie
}

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

func (svc *AuthService) UpdateUserPassword(ctx context.Context, token, newPassword string) error {
	if newPassword == "" {
		return fmt.Errorf("New password is missing")
	}

	email := jwtoken.GetUserEmail(token)
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

	return nil
}
