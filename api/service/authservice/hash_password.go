package authservice

import "golang.org/x/crypto/bcrypt"

const passwordCost = bcrypt.DefaultCost + 2

func (svc *AuthService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
