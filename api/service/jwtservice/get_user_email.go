package jwtservice

func GetUserEmail(tokenString string) string {
	claims := getClaims(tokenString)

	if claims != nil {
		return claims.Email
	}

	return ""
}
