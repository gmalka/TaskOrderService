package ordercontroller

import (
	"taskServer/database"
	"taskServer/model"
	"taskServer/nosql"
)

type Controller interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	GetTask(id int) (model.Task, error)
	ChangeTaskPrice(id int, price int) error
	DeleteTask(id int) error
	GetAllTasksOfUser(username string, page int) ([]model.Task, error)
	DeleteAllTasksOfUser(username string) error
	BuyTaskAnswer(task model.UsersPurchase) error
}

type orderController struct {
	db    database.DatabaseService
	cache nosql.NoSqlService
}

func NewUserController(db database.DatabaseService) Controller {
	return orderController{
		db: db,
	}
}

func (o orderController) GetAllTasks() ([]model.Task, error) {
	return o.db.GetAllTasks()
}

func (o orderController) CreateTask(task model.Task) error {
	return o.db.CreateTask(task)
}

func (o orderController) GetTask(id int) (model.Task, error) {
	return o.db.GetTask(id)
}

func (o orderController) ChangeTaskPrice(id int, price int) error {
	return o.db.ChangeTaskPrice(id, price)
}

func (o orderController) DeleteTask(id int) error {
	return o.db.DeleteTask(id)
}

func (o orderController) GetAllTasksOfUser(username string, page int) ([]model.Task, error) {
	return o.db.GetAllTasksOfUser(username, page)
}

func (o orderController) DeleteAllTasksOfUser(username string) error {
	return o.db.DeleteAllTasksOfUser(username)
}

func (o orderController) BuyTaskAnswer(task model.UsersPurchase) error {
	return o.db.BuyTaskAnswer(task)
}
