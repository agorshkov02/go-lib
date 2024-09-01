package bcrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hash, password string) (bool, error) {
	hashBytes := []byte(hash)
	passwordBytes := []byte(password)
	if err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GenerateFromPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashBytes, err := bcrypt.GenerateFromPassword(passwordBytes, 12)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}
