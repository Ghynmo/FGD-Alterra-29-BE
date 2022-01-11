package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(secret string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ValidateHash(secret string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
