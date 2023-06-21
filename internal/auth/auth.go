package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing password")
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(verifiable, wanted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(wanted), []byte(verifiable))
	if err != nil {
		log.Println("Error: failed authorization")
		return  false
	}

	return true
}