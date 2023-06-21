package database

import "userService/internal/model"


type DatabaseService interface {
	Create(user model.UserWithRole) error
	GetByUsername(username string) (model.UserInfoWithRoleruct, error)
	GetAllUsers() ([]string, error)
	Update(model.User) error
	Delete(string) error
}