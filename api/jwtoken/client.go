package jwtoken

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type JwtClient struct {
	claims JwtClaims
}

const ExpirationDuration = time.Hour * (24 * 5)

func New(userId, email string) *JwtClient {
	now := time.Now()

	client := &JwtClient{
		claims: JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    os.Getenv("JWT_ISS"),
				Subject:   userId,
				Audience:  jwt.ClaimStrings{os.Getenv("JWT_AUD")},
				ExpiresAt: jwt.NewNumericDate(now.Add(ExpirationDuration)),
				NotBefore: jwt.NewNumericDate(now),
				IssuedAt:  jwt.NewNumericDate(now),
				ID:        uuid.NewString(),
			},
			Email: email,
		},
	}

	return client
}
