package jwtoken

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JwtClient) GenerateToken() (string, error) {
	if j.claims.RegisteredClaims.Subject == "" || j.claims.Email == "" {
		return "", fmt.Errorf("token subject and email cannot be empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenString(authHeader string) (string, error) {
	if authHeader == "" {
		return "", nil
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			return "", fmt.Errorf("token is missing")
		}

		return tokenString, nil
	}

	return "", fmt.Errorf("invalid authorization header format")
}

func getClaims(tokenString string) *JwtClaims {
	if tokenString == "" {
		return nil
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatalf("JWT secret is missing")
		return nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		slog.Info(fmt.Sprintf("Error parsing token: %v", err.Error()))
		return nil
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims
	}

	return nil
}

func GetUserId(tokenString string) string {
	claims := getClaims(tokenString)

	if claims != nil {
		return claims.Subject
	}

	return ""
}

func GetUserEmail(tokenString string) string {
	claims := getClaims(tokenString)

	if claims != nil {
		return claims.Email
	}

	return ""
}

func VerifyToken(tokenString string) bool {
	claims := getClaims(tokenString)

	return claims != nil
}
