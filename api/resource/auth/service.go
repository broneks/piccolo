package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const COST = bcrypt.DefaultCost + 2

func setAccessTokenCookie(c echo.Context, value string) {
	c.SetCookie(&http.Cookie{
		Name:     "piccolo-access-token",
		Value:    value,
		HttpOnly: true,
		Secure:   false, // TODO change this for production
		// SameSite: http.SameSiteStrictMode, // Prevents CSRF by restricting cross-site cookie transmission TODO
		Path: "/",
	})
}

func setRefreshTokenCookie(c echo.Context, value string) {
	c.SetCookie(&http.Cookie{
		Name:     "piccolo-refresh-token",
		Value:    value,
		HttpOnly: true,
		Secure:   false, // TODO change this for production
		// SameSite: http.SameSiteStrictMode, // Prevents CSRF by restricting cross-site cookie transmission TODO
		Path: "/",
	})
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
