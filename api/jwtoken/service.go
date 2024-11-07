package jwtoken

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JwtClient) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenString(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
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
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return os.Getenv("JWT_SECRET"), nil
	})
	if err != nil {
		log.Fatalf("Error parsing token: %v", err)
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

	if claims != nil {
		fmt.Printf("Token is valid! Email: %s\n", claims.Email)
		return true
	}

	log.Println("Token is invalid or expired.")
	return false
}
