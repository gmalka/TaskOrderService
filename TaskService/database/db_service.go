package database

import "taskServer/model"


type DatabaseService interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	GetTask(id int) (model.Task, error)
	ChangeTaskPrice(id int, price int) error
	DeleteTask(id int) error
	GetAllTasksOfUser(username string, page int) ([]model.Task, error)
	DeleteAllTasksOfUser(username string) error
	BuyTaskAnswer(task model.UsersPurchase) error
}