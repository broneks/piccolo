package authservice

import (
	"net/http"
	"os"
	"time"
)

func (svc *AuthService) NewAccessTokenCookie(value string) *http.Cookie {
	env := os.Getenv("ENV")

	cookie := &http.Cookie{
		Name:     "piccolo-access-token",
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	if env == "local" {
		cookie.Secure = false
		cookie.SameSite = http.SameSiteNoneMode
	}

	if value == "" {
		cookie.Expires = time.Unix(0, 0)
		cookie.MaxAge = -1
	}

	return cookie
}
