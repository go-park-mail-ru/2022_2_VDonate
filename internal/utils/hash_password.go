package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 14

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", nil
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
