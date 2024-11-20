package jwtoken

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

type JwtClient struct {
	claims JwtClaims
}

const AccessExpirationDuration = time.Hour * 4          // 4 hours
const RefreshExpirationDuration = time.Hour * (24 * 14) // 14 days
const ResetPasswordExpirationDuration = time.Hour       // 1 hour

func New(action, subject, email string, expirationDuration time.Duration) *JwtClient {
	now := time.Now()

	client := &JwtClient{
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
			Email: email,
			Action: action,
		},
	}

	return client
}

func NewAccessJwt(userId, email string) *JwtClient {
	return New("access", userId, email, AccessExpirationDuration)
}

func NewRefreshJwt(userId, email string) *JwtClient {
	return New("refresh", userId, email, RefreshExpirationDuration)
}

func NewResetPasswordJwt(email string) *JwtClient {
	return New("reset-password", email, email, ResetPasswordExpirationDuration)
}
