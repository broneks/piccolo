package authservice

import (
	"net/http"
	"time"
)

func (svc *AuthService) NewAccessTokenCookie(value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "piccolo-access-token",
		Value:    value,
		HttpOnly: true,
		Secure:   false, // TODO change this for production
		// SameSite: http.SameSiteStrictMode, // Prevents CSRF by restricting cross-site cookie transmission TODO
		Path: "/",
	}

	if value == "" {
		cookie.Expires = time.Unix(0, 0)
	}

	return cookie
}
