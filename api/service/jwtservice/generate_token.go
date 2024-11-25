package jwtservice

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func (jc *JwtService) GenerateToken() (string, error) {
	if jc.claims.RegisteredClaims.Subject == "" || jc.claims.Email == "" {
		return "", fmt.Errorf("token subject and email cannot be empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc.claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
