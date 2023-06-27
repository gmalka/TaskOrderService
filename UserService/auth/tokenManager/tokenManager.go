package tokenManager

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	ACCESS_TOKEN_TTL  = 15
	REFRESH_TOKEN_TTL = 60

	AccessToken = iota
	RefreshToken
)

type UserClaims struct {
	Username  string `json:"username"`
	Role      string `json:"role"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	jwt.RegisteredClaims
}

type UserInfo struct {
	Username  string `json:"username"`
	Role      string `json:"role"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type authService struct {
	accessSecret  []byte
	refreshSecret []byte
}

type TokenManager interface {
	CreateToken(userinfo UserInfo, ttl time.Duration, kind int) (string, error)
	ParseToken(inputToken string, kind int) (UserClaims, error)
}

func NewAuthService(accessSecret, refreshSecret string) TokenManager {
	return authService{
		accessSecret:  []byte(accessSecret),
		refreshSecret: []byte(refreshSecret),
	}
}

func (u authService) ParseToken(inputToken string, kind int) (UserClaims, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		var secret []byte
		switch kind {
		case AccessToken:
			secret = u.accessSecret
		case RefreshToken:
			secret = u.refreshSecret
		default:
			return "", fmt.Errorf("unknown secret kind %d", kind)
		}

		return secret, nil
	})

	if err != nil {
		return UserClaims{}, fmt.Errorf("can't parse token: %v", err)
	}

	if !token.Valid {
		return UserClaims{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserClaims{}, fmt.Errorf("can't get user claims from token")
	}

	return UserClaims{
		Username:         claims["username"].(string),
		Role:             claims["role"].(string),
		Firstname:        claims["firstname"].(string),
		Lastname:         claims["lastname"].(string),
		RegisteredClaims: jwt.RegisteredClaims{},
	}, nil
}

func (u authService) CreateToken(userinfo UserInfo, ttl time.Duration, kind int) (string, error) {
	claims := UserClaims{
		Username:         userinfo.Username,
		Role:             userinfo.Role,
		Firstname:        userinfo.Firstname,
		Lastname:         userinfo.Lastname,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl * time.Minute))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var secret []byte
	switch kind {
	case AccessToken:
		secret = u.accessSecret
	case RefreshToken:
		secret = u.refreshSecret
	default:
		return "", fmt.Errorf("unknown secret kind %d", kind)
	}

	return token.SignedString(secret)
}
