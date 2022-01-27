package helpers

import "golang.org/x/crypto/bcrypt"

type Bycript interface {
	Hash(password string) (string, error)
	ValidateHash(password, hash string) bool
}

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
