package jwtservice

func GetUserId(tokenString string) string {
	claims := getClaims(tokenString)

	if claims != nil {
		return claims.Subject
	}

	return ""
}
