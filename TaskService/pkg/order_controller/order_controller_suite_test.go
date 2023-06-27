package ordercontroller_test

import (
	"TaskService/pkg/model"
	"testing"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

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

type DbDouble struct {
	Double
}

func NewDbDouble() DbDouble {
	return DbDouble{Double: NewStrictDouble()}
}

func TestOrderController(t *testing.T) {
	RegisterDoublesFailHandler(func(message string, callerSkip ...int) {
		t.Fatal(message)
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "OrderController Suite")
}

func (d DbDouble) GetAllTasks() ([]model.Task, error) {
	returnVal, _ := d.Call("GetAllTasks")
	rFirst, _ := returnVal[0].([]model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d DbDouble) CreateTask(task model.Task) error {
	returnVal, _ := d.Call("CreateTask", task)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d DbDouble) CheckAndGetTask(username string, id int) (model.Task, error) {
	returnVal, _ := d.Call("CheckAndGetTask", username, id)
	rFirst, _ := returnVal[0].(model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d DbDouble) ChangeTaskPrice(id int, price int) error {
	returnVal, _ := d.Call("ChangeTaskPrice", id, price)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d DbDouble) DeleteTask(id int) error {
	returnVal, _ := d.Call("DeleteTask", id)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d DbDouble) GetAllTasksOfUser(username string, page int) ([]model.Task, error) {
	returnVal, _ := d.Call("GetAllTasksOfUser", username, page)
	rFirst, _ := returnVal[0].([]model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d DbDouble) GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error) {
	returnVal, _ := d.Call("GetAllTasksWithoutAnswers", page)
	rFirst, _ := returnVal[0].([]model.TaskWithoutAnswer)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d DbDouble) DeleteAllTasksOfUser(username string) error {
	returnVal, _ := d.Call("DeleteAllTasksOfUser", username)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d DbDouble) BuyTaskAnswer(task model.UsersPurchase) error {
	returnVal, _ := d.Call("BuyTaskAnswer", task)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}
