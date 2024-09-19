package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return errors.New("invalid email and/or password")
	}
	return nil
}
