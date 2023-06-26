package rest_test

import (
	"taskServer/model"
	"testing"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

func TestRest(t *testing.T) {
	RegisterDoublesFailHandler(func(message string, callerSkip ...int) {
		t.Fatal(message)
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest Suite")
}

type ControllerDouble struct {
	Double
}

func NewDbDouble() ControllerDouble {
	return ControllerDouble{Double: NewStrictDouble()}
}


func (d ControllerDouble)GetAllTasks() ([]model.Task, error){
	returnVal, _ := d.Call("GetAllTasks")
	rFirst, _ := returnVal[0].([]model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d ControllerDouble)CreateTask(task model.Task) error{
	returnVal, _ := d.Call("CreateTask", task)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d ControllerDouble)CheckAndGetTask(username string, id int) (model.Task, error){
	returnVal, _ := d.Call("CheckAndGetTask", username, id)
	rFirst, _ := returnVal[0].(model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d ControllerDouble)ChangeTaskPrice(id int, price int) error{
	returnVal, _ := d.Call("ChangeTaskPrice", id, price)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d ControllerDouble)DeleteTask(id int) error{
	returnVal, _ := d.Call("DeleteTask", id)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d ControllerDouble)GetAllTasksOfUser(username string, page int) ([]model.Task, error){
	returnVal, _ := d.Call("GetAllTasksOfUser", username, page)
	rFirst, _ := returnVal[0].([]model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d ControllerDouble)GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error){
	returnVal, _ := d.Call("GetAllTasksWithoutAnswers", page)
	rFirst, _ := returnVal[0].([]model.TaskWithoutAnswer)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (d ControllerDouble)DeleteAllTasksOfUser(username string) error{
	returnVal, _ := d.Call("DeleteAllTasksOfUser", username)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (d ControllerDouble) BuyTaskAnswer(task model.UsersPurchase) error{
	returnVal, _ := d.Call("BuyTaskAnswer", task)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}