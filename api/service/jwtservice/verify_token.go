package jwtservice

func VerifyToken(tokenString string) bool {
	claims := getClaims(tokenString)

	return claims != nil
}
