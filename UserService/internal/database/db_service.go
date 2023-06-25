package database

import "userService/internal/model"


type DatabaseService interface {
	Create(user model.UserWithRole) error
	GetByUsername(username string) (model.UserWithRole, error)
	GetAllUsers() ([]string, error)
	TryToBuyTask(username string, price int) (error)
	Update(user model.UserForUpdate) error
	Delete(username string) error
	UpdateBalance(username string, change int) error
}