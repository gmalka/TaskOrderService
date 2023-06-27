package rest_test

import (
	"testing"
	"time"
	"userService/auth/tokenManager"
	"userService/pkg/model"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRest(t *testing.T) {
	RegisterDoublesFailHandler(func(message string, callerSkip ...int) {
		t.Fatal(message)
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest Suite")
}

// ==================BD=MOK====================> //

type UserControllerDouble struct {
	Double
}

func NewUserControllerDouble() UserControllerDouble {
	return UserControllerDouble{Double: NewStrictDouble()}
}

func (uc UserControllerDouble) CreateUser(user model.User) error {
	returnVal, _ := uc.Call("CreateUser", user)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UserControllerDouble) GetUser(username string) (model.UserWithRole, error) {
	returnVal, _ := uc.Call("GetUser", username)
	rFirst, _ := returnVal[0].(model.UserWithRole)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (uc UserControllerDouble) GetAllUsernames() ([]string, error) {
	returnVal, _ := uc.Call("GetAllUsernames")
	rFirst, _ := returnVal[0].([]string)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (uc UserControllerDouble) TryToBuyTask(username string, price int) error {
	returnVal, _ := uc.Call("TryToBuyTask", username, price)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UserControllerDouble) UpdateUser(user model.UserForUpdate) error {
	returnVal, _ := uc.Call("UpdateUser", user)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UserControllerDouble) DeleteUser(username string) error {
	returnVal, _ := uc.Call("DeleteUser", username)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UserControllerDouble) UpdateBalance(username string, change int) error {
	returnVal, _ := uc.Call("UpdateBalance", username, change)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

// ============================================> //

// ================GRPC=MOK====================> //
type GrpcServiceDouble struct {
	Double
}

func NewGrpcServiceDouble() GrpcServiceDouble {
	return GrpcServiceDouble{Double: NewStrictDouble()}
}

func (g GrpcServiceDouble) UpdatePriceOfTask(id, price int) error {
	returnVal, _ := g.Call("UpdatePriceOfTask", id, price)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (g GrpcServiceDouble) CreateNewTask(task model.TaskWithoutAnswer) error {
	returnVal, _ := g.Call("CreateNewTask", task)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (g GrpcServiceDouble) CheckAndGetTask(username string, id int) (model.TaskOrderInfo, error) {
	returnVal, _ := g.Call("CheckAndGetTask", username, id)
	rFirst, _ := returnVal[0].(model.TaskOrderInfo)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (g GrpcServiceDouble) BuyTaskAnswer(username string, taskId int) error {
	returnVal, _ := g.Call("BuyTaskAnswer", username, taskId)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (g GrpcServiceDouble) GetAllTasks() ([]model.Task, error) {
	returnVal, _ := g.Call("GetAllTasks")
	rFirst, _ := returnVal[0].([]model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (g GrpcServiceDouble) GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error) {
	returnVal, _ := g.Call("GetAllTasksWithoutAnswers", page)
	rFirst, _ := returnVal[0].([]model.TaskWithoutAnswer)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (g GrpcServiceDouble) GetOrdersForUser(username string, page int) ([]model.Task, error) {
	returnVal, _ := g.Call("GetOrdersForUser", username, page)
	rFirst, _ := returnVal[0].([]model.Task)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (g GrpcServiceDouble) DeleteOrdersForUser(username string) error {
	returnVal, _ := g.Call("DeleteOrdersForUser", username)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (g GrpcServiceDouble) DeleteTask(taskId int) error {
	returnVal, _ := g.Call("DeleteTask", taskId)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

// ============================================> //

// ===============TOKEN=MOK====================> //
type TokenManagerDouble struct {
	Double
}

func NewTokenManagerDouble() TokenManagerDouble {
	return TokenManagerDouble{Double: NewStrictDouble()}
}

func (t TokenManagerDouble) CreateToken(userinfo tokenManager.UserInfo, ttl time.Duration, kind int) (string, error) {
	returnVal, _ := t.Call("CreateToken", userinfo, ttl, kind)
	rFirst, _ := returnVal[0].(string)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (t TokenManagerDouble) ParseToken(inputToken string, kind int) (tokenManager.UserClaims, error) {
	returnVal, _ := t.Call("ParseToken", inputToken, kind)
	rFirst, _ := returnVal[0].(tokenManager.UserClaims)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

// ============================================> //

// =================PAS=MOK====================> //
type PasswordMangerDouble struct {
	Double
}

func NewPasswordMangerDouble() PasswordMangerDouble {
	return PasswordMangerDouble{Double: NewStrictDouble()}
}

func (p PasswordMangerDouble) HashPassword(password string) (string, error) {
	returnVal, _ := p.Call("HashPassword", password)
	rFirst, _ := returnVal[0].(string)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (p PasswordMangerDouble) CheckPassword(verifiable, wanted string) error {
	returnVal, _ := p.Call("CheckPassword", verifiable, wanted)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

// ============================================> //
