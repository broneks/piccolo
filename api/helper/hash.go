package helper

import (
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/sha3"
)

func generateRandomSalt(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	salt := make([]byte, length)
	for i := 0; i < length; i++ {
		salt[i] = chars[r.Intn(len(chars))]
	}

	return string(salt)
}

func GenerateRandomHash() string {
	salt := generateRandomSalt(42)

	hash := sha3.New256()
	hash.Write([]byte(salt))

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
