package usercontroller

import (
	"userService/internal/database"
	"userService/internal/model"
)

const (
	REGULAR_ROLE = "regular"
	ADMIN_ROLE   = "admin"
)

type Controller interface {
	CreateUser(user model.User) error
	GetAllUsernames() ([]string, error)
	GetUser(username string) (model.UserWithRole, error)
	TryToBuyTask(username string, price int) (error)
	DeleteUser(username string) error
	UpdateUser(user model.UserForUpdate) error
	UpdateBalance(username string, change int) error
}

type userController struct {
	db database.DatabaseService
}

func NewUserController(db database.DatabaseService) Controller {
	return userController{
		db: db,
	}
}

func (u userController) UpdateBalance(username string, change int) error {
	return u.db.UpdateBalance(username, change)
}

func (u userController) CreateUser(user model.User) error {
	return u.db.Create(model.UserWithRole{User: user, Role: REGULAR_ROLE})
}

func (u userController) UpdateUser(user model.UserForUpdate) error {
	return u.db.Update(user)
}

func (u userController) GetAllUsernames() ([]string, error) {
	return u.db.GetAllUsers()
}

func (u userController) GetUser(username string) (model.UserWithRole, error) {
	return u.db.GetByUsername(username)
}

func (u userController) TryToBuyTask(username string, price int) (error) {
	return u.db.TryToBuyTask(username, price)
}

func (u userController) DeleteUser(username string) error {
	return u.db.Delete(username)
}