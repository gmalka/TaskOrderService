package database

import "TaskService/pkg/model"

type DatabaseService interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	CheckAndGetTask(username string, id int) (model.Task, error)
	ChangeTaskPrice(id int, price int) error
	DeleteTask(id int) error
	GetAllTasksOfUser(username string, page int) ([]model.Task, error)
	GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error)
	DeleteAllTasksOfUser(username string) error
	BuyTaskAnswer(task model.UsersPurchase) error
}
