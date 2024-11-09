package jwtoken

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JwtClient struct {
	claims JwtClaims
}

const AccessExpirationDuration = time.Hour              // 1 hour
const RefreshExpirationDuration = time.Hour * (24 * 14) // 14 days

func New(userId, email string, expirationDuration time.Duration) *JwtClient {
	now := time.Now()

	client := &JwtClient{
		claims: JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    os.Getenv("JWT_ISS"),
				Subject:   userId,
				Audience:  jwt.ClaimStrings{os.Getenv("JWT_AUD")},
				ExpiresAt: jwt.NewNumericDate(now.Add(expirationDuration)),
				NotBefore: jwt.NewNumericDate(now),
				IssuedAt:  jwt.NewNumericDate(now),
				ID:        uuid.NewString(),
			},
			Email: email,
		},
	}

	return client
}

func NewAccessJwt(userId, email string) *JwtClient {
	return New(userId, email, AccessExpirationDuration)
}

func NewRefreshJwt(userId, email string) *JwtClient {
	return New(userId, email, RefreshExpirationDuration)
}
