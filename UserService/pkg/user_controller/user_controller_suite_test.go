package usercontroller_test

import (
	"testing"
	"userService/pkg/model"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserController(t *testing.T) {
	RegisterDoublesFailHandler(func(message string, callerSkip ...int) {
		t.Fatal(message)
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserController Suite")
}

type UcDouble struct {
	Double
}

func NewUcDouble() UcDouble {
	return UcDouble{Double: NewStrictDouble()}
}

func (uc UcDouble) Create(user model.UserWithRole) error {
	returnVal, _ := uc.Call("Create", user)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UcDouble) GetByUsername(username string) (model.UserWithRole, error) {
	returnVal, _ := uc.Call("GetByUsername", username)
	rSecond, _ := returnVal[0].(model.UserWithRole)
	rFirst, _ := returnVal[1].(error)
	return rSecond, rFirst
}

func (uc UcDouble) GetAllUsers() ([]string, error) {
	returnVal, _ := uc.Call("GetAllUsers")
	rFirst, _ := returnVal[0].([]string)
	rSecond, _ := returnVal[1].(error)
	return rFirst, rSecond
}

func (uc UcDouble) TryToBuyTask(username string, price int) error {
	returnVal, _ := uc.Call("TryToBuyTask", username, price)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UcDouble) Update(user model.UserForUpdate) error {
	returnVal, _ := uc.Call("Update", user)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UcDouble) Delete(username string) error {
	returnVal, _ := uc.Call("Delete", username)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}

func (uc UcDouble) UpdateBalance(username string, change int) error {
	returnVal, _ := uc.Call("UpdateBalance", username, change)
	rFirst, _ := returnVal[0].(error)
	return rFirst
}
