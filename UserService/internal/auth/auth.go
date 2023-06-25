package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("can't hash password: %v", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(verifiable, wanted string) error {
	err := bcrypt.CompareHashAndPassword([]byte(wanted), []byte(verifiable))
	if err != nil {
		return fmt.Errorf("authorization failed: %v", err)
	}

	return nil
}