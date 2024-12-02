package jwtservice

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func getClaims(tokenString string) *JwtClaims {
	if tokenString == "" {
		return nil
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatalf("JWT secret is missing")
		return nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		slog.Debug("error parsing token", "err", err)
		return nil
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims
	}

	return nil
}
