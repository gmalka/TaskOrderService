package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessToken = iota
	RefreshToken
)

type authService struct {
	accessSecret  []byte
	refreshSecret []byte
}

type TokenManager interface {
	CreateToken(role, username string, ttl time.Duration, kind int) (string, error)
	ParseToken(inputToken string, kind int) (UserClaims, error)
}

type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (u authService) ParseToken(inputToken string, kind int) (UserClaims, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		var secret []byte
		switch kind {
		case AccessToken:
			secret = u.accessSecret
		case RefreshToken:
			secret = u.refreshSecret
		default:
			return "", fmt.Errorf("unknown secret kind: %d", kind)
		}

		return secret, nil
	})

	if err != nil {
		return UserClaims{}, err
	}

	if !token.Valid {
		return UserClaims{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserClaims{}, fmt.Errorf("error get user claims from token")
	}


	return UserClaims{
		Username: claims["username"].(string),
		Role: claims["role"].(string),
		RegisteredClaims: jwt.RegisteredClaims{},
	 }, nil
}

func (u authService) CreateToken(role, username string, ttl time.Duration, kind int) (string, error) {
	claims := UserClaims{
		Username:         username,
		Role:             role,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var secret []byte
	switch kind {
	case AccessToken:
		secret = u.accessSecret
	case RefreshToken:
		secret = u.refreshSecret
	default:
		return "", fmt.Errorf("unknown secret kind: %d", kind)
	}

	return token.SignedString(secret)
}
