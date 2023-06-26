package passwordHandler

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHandler interface {
	HashPassword(password string) (string, error)
	CheckPassword(verifiable, wanted string) error
}

type passwordHandler struct {
}

func NewPasswordManager() PasswordHandler {
	return passwordHandler{}
}

func (p passwordHandler) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("can't hash password: %v", err)
	}

	return string(hashedPassword), nil
}

func (p passwordHandler) CheckPassword(verifiable, wanted string) error {
	err := bcrypt.CompareHashAndPassword([]byte(wanted), []byte(verifiable))
	if err != nil {
		return fmt.Errorf("authorization failed: %v", err)
	}

	return nil
}