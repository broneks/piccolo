package helper

import (
	"fmt"
	"strings"
)

func ExtractTokenString(authHeader string) (string, error) {
	if authHeader == "" {
		return "", nil
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			return "", fmt.Errorf("token is missing")
		}

	}

	return "", fmt.Errorf("invalid authorization header format")
}
