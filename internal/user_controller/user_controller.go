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
		nosql: nosql,
		db: db,
		tokenManager: tokenManager,
	}
}

func (u userController) GetUser(userauth model.UserAuth) (model.UserWithRole, error) {
	
}

func (u userController) RegisterUser(user model.User) error {
	err := u.db.Create(model.UserWithRole{
		User: user,
		Role: "user",
	})

	return err
}