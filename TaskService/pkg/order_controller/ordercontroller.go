package ordercontroller

import (
	"TaskService/pkg/database"
	"TaskService/pkg/model"
)

type Controller interface {
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

type orderController struct {
	db database.DatabaseService
}

func NewUserController(db database.DatabaseService) Controller {
	return orderController{
		db: db,
	}
}

func (o orderController) GetAllTasks() ([]model.Task, error) {
	return o.db.GetAllTasks()
}

func (o orderController) GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error) {
	return o.db.GetAllTasksWithoutAnswers(page)
}

func (o orderController) CreateTask(task model.Task) error {
	task.Answer = calculateAnswer(task.Heights)
	return o.db.CreateTask(task)
}

func calculateAnswer(mas []int64) int {
	max := 0;
	cur := 0
	sum := 0

	if len(mas) < 2 {
		return 0
	}

	for i := 0; i < len(mas); i++ {
		if mas[i] > mas[max] {
			max = i
		}
	}

	for i := 0; i < max; i++ {
		if mas[cur] < mas[i] {
			cur = i 
		} else {
			sum += int(mas[cur] - mas[i])
		}
	}

	cur = len(mas) - 1
	for i := len(mas) - 1; i > max; i-- {
		if mas[cur] <= mas[i] {
			cur = i 
		} else {
			sum += int(mas[cur] - mas[i])
		}
	}

	return sum
}

func (o orderController) CheckAndGetTask(username string, id int) (model.Task, error) {
	return o.db.CheckAndGetTask(username, id)
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
