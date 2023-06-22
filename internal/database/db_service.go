package database

import "userService/internal/model"


type DatabaseService interface {
	Create(user model.UserWithRole) error
	GetByUsername(username string) (model.UserWithRole, error)
	GetAllUsers() ([]string, error)
	GetOrdersOfUser(username string, number int) ([]model.Task, error)
	TryToBuyTask(username string, price int) (bool, error)
	Update(model.UserForUpdate) error
	Delete(string) error
}