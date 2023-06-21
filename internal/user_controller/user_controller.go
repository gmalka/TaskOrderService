package usercontroller

import (
	"userService/internal/auth"
	"userService/internal/database"
	"userService/internal/model"
	"userService/internal/nosql"
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

func (u userController) GetAllUsernames(username string) ([]string, error) {
	return u.db.GetAllUsers()
}

func (u userController) GetUserInfo(username string) (model.UserInfoWithRole, error) {
	return u.db.GetByUsername(username)
}

func (u userController) RegisterUser(user model.User) error {
	err := u.db.Create(model.UserWithRole{
		User: user,
		Role: "user",
	})

	return err
}

func (u userController) LoginUser(userAuth model.UserAuth) (string, string, error) {
	password, err := u.db.GetPassword(userAuth.Username)
	if err != nil {
		return "", "", err
	}

}
