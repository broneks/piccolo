package jwtservice

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Email  string `json:"email"`
	Action string `json:"action"`
}

type JwtService struct {
	claims JwtClaims
}

const accessExpirationDuration = time.Minute * 15       // 15 minutes
const refreshExpirationDuration = time.Hour * (24 * 14) // 14 days
const resetPasswordExpirationDuration = time.Hour       // 1 hour

func New(action, subject, email string, expirationDuration time.Duration) *JwtService {
	now := time.Now()

	client := &JwtService{
		claims: JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    os.Getenv("JWT_ISS"),
				Subject:   subject,
				Audience:  jwt.ClaimStrings{os.Getenv("JWT_AUD")},
				ExpiresAt: jwt.NewNumericDate(now.Add(expirationDuration)),
				NotBefore: jwt.NewNumericDate(now),
				IssuedAt:  jwt.NewNumericDate(now),
				ID:        uuid.NewString(),
			},
			Email:  email,
			Action: action,
		},
	}

	return client
}

func NewAccessJwt(userId, email string) *JwtService {
	return New("access", userId, email, accessExpirationDuration)
}

func NewRefreshJwt(userId, email string) *JwtService {
	return New("refresh", userId, email, refreshExpirationDuration)
}

func NewResetPasswordJwt(email string) *JwtService {
	return New("reset-password", email, email, resetPasswordExpirationDuration)
}
