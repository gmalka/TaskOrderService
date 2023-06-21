package usercontroller

import (
	"errors"
	"fmt"
	"userService/internal/auth"
	"userService/internal/database"
	"userService/internal/model"
	"userService/internal/nosql"
)

const (
	ACCESS_TOKEN_TTL=15
	REFRESH_TOKEN_TTL=60
)

const (
	// key should be like `username:key`
	ACCESS = "access"
	REFRESH = "refresh"
)

type Controller interface {
}

type userController struct {
	nosql        nosql.NoSqlService
	db           database.DatabaseService
	tokenManager auth.TokenManager
}

func NewUserController(nosql nosql.NoSqlService, db database.DatabaseService, tokenManager auth.TokenManager) Controller {
	return userController{
		nosql:        nosql,
		db:           db,
		tokenManager: tokenManager,
	}
}

// as kind must be used usercontroller.ACCESS or usercontroller.REFRESH
func (u userController) GetUserToken(username, kind string) (string, error) {
	return u.nosql.Get(fmt.Sprintf("%s:%s", username, kind))
}

func (u userController) GetAllUsernames(username string) ([]string, error) {
	return u.db.GetAllUsers()
}

func (u userController) GetUserInfo(username string) (model.UserInfo, error) {
	user, err := u.db.GetByUsername(username)
	return user.User.Info, err
}

func (u userController) DeleteUser(username string) error {
	err := u.db.Delete(username)
	if err != nil {
		return err
	}

	u.nosql.Delete(username)
	return nil
}

func (u userController) RegisterUser(user model.User) error {
	var err error

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = u.db.Create(model.UserWithRole{
		User: user,
		Role: "user",
	})

	return err
}

func (u userController) LoginUser(userAuth model.UserAuth) (string, string, error) {
	user, err := u.db.GetByUsername(userAuth.Username)
	if err != nil {
		return "", "", err
	}

	if ok := auth.CheckPassword(userAuth.Password, user.User.Password); !ok {
		return "", "", errors.New("passwords mismatch")
	}

	accessToken, err := u.tokenManager.CreateToken(auth.UserInfo{
		Username: user.User.Username,
		Role: user.Role,
		Firstname: user.User.Info.Firstname,
		Lastname: user.User.Info.Lastname,
	}, ACCESS_TOKEN_TTL, auth.AccessToken)
	if err != nil {
		return "", "", err
	}

	err = u.nosql.Set(fmt.Sprintf("%s:%s", user.User.Username, ACCESS), accessToken, ACCESS_TOKEN_TTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.tokenManager.CreateToken(auth.UserInfo{
		Username: user.User.Username,
		Role: user.Role,
		Firstname: user.User.Info.Firstname,
		Lastname: user.User.Info.Lastname,
	}, REFRESH_TOKEN_TTL, auth.RefreshToken)
	if err != nil {
		return "", "", err
	}

	err = u.nosql.Set(fmt.Sprintf("%s:%s", user.User.Username, REFRESH), refreshToken, REFRESH_TOKEN_TTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}