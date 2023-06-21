package database

import "userService/internal/model"


type DatabaseService interface {
	Create(user model.UserWithRole) error
	GetByUsername(username string) (model.UserWithRole, error)
	GetPassword(username string) (string, error)
	GetAllUsers() ([]string, error)
	Update(model.User) error
	Delete(string) error
}