package database

import "userService/internal/model"


type DatabaseService interface {
	Create(user model.UserWithRole) error
	GetByUsername(username string) (model.UserWithRole, error)
	GetAllUsers() ([]string, error)
	Update(model.UserForUpdate) error
	Delete(string) error
}