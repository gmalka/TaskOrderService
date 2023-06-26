package usercontroller_test

import (
	"errors"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"userService/internal/model"
	usercontroller "userService/internal/user_controller"
)

var _ = Describe("UserController", func() {
	var (
		uc UcDouble
		c  usercontroller.Controller
	)

	BeforeEach(func() {
		uc = NewUcDouble()
		c = usercontroller.NewUserController(uc)
	})

	Context("user controller test", func() {
		Context("testing CreateUser", func() {
			It("regular", func() {
				user := model.User{
					Username: "root",
					Password: "1234",
					Info: model.UserInfo{
						Firstname: "m",
						Lastname:  "m",
						Surname:   "mm",
						Group:     "p-14",
						Balance:   1500,
					},
				}
				uWithRole := model.UserWithRole{
					User: user,
					Role: "regular",
				}
				AllowDouble(uc).To(ReceiveCallTo("Create").With(uWithRole).AndReturn(nil))

				Expect(c.CreateUser(user)).Should(Succeed())
			})
			It("error", func() {
				user := model.User{
					Username: "root",
					Password: "1234",
					Info: model.UserInfo{
						Firstname: "m",
						Lastname:  "m",
						Surname:   "mm",
						Group:     "p-14",
						Balance:   1500,
					},
				}
				uWithRole := model.UserWithRole{
					User: user,
					Role: "regular",
				}
				AllowDouble(uc).To(ReceiveCallTo("Create").With(uWithRole).AndReturn(errors.New("some error")))

				Expect(c.CreateUser(user)).ShouldNot(Succeed())
			})
		})

		Context("testing GetAllUsernames", func() {
			It("regular", func() {
				want := []string{"root", "admin", "gmalka"}
				AllowDouble(uc).To(ReceiveCallTo("GetAllUsers").With().AndReturn(want, nil))

				Expect(c.GetAllUsernames()).To(Equal(want))
			})
			It("error", func() {
				AllowDouble(uc).To(ReceiveCallTo("GetAllUsers").With().AndReturn(nil, errors.New("some error")))

				_, err := c.GetAllUsernames()
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing GetUser", func() {
			It("regular", func() {
				user := model.User{
					Username: "root",
					Password: "1234",
					Info: model.UserInfo{
						Firstname: "m",
						Lastname:  "m",
						Surname:   "mm",
						Group:     "p-14",
						Balance:   1500,
					},
				}
				uWithRole := model.UserWithRole{
					User: user,
					Role: "regular",
				}
				AllowDouble(uc).To(ReceiveCallTo("GetByUsername").With("root").AndReturn(uWithRole, nil))

				Expect(c.GetUser("root")).To(Equal(uWithRole))
			})
			It("error", func() {
				AllowDouble(uc).To(ReceiveCallTo("GetByUsername").With("root").AndReturn(nil, errors.New("some error")))

				_, err := c.GetUser("root")
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing TryToBuyTask", func() {
			It("regular", func() {
				AllowDouble(uc).To(ReceiveCallTo("TryToBuyTask").With("root", 1).AndReturn(nil))

				Expect(c.TryToBuyTask("root", 1)).Should(Succeed())
			})
			It("error", func() {
				AllowDouble(uc).To(ReceiveCallTo("TryToBuyTask").With("root", 1).AndReturn(errors.New("some error")))

				Expect(c.TryToBuyTask("root", 1)).ShouldNot(Succeed())
			})
		})

		Context("testing DeleteUser", func() {
			It("regular", func() {
				AllowDouble(uc).To(ReceiveCallTo("Delete").With("root").AndReturn(nil))

				Expect(c.DeleteUser("root")).Should(Succeed())
			})
			It("error", func() {
				AllowDouble(uc).To(ReceiveCallTo("Delete").With("root").AndReturn(errors.New("some error")))

				Expect(c.DeleteUser("root")).ShouldNot(Succeed())
			})
		})

		Context("testing UpdateUser", func() {
			It("regular", func() {
				user := model.UserForUpdate{
					Username: "root",
					Password: "1234",
					Info: model.UserInfoForUpdate{
						Firstname: "m",
						Lastname:  "m",
						Surname:   "mm",
						Group:     "p-14",
					},
				}
				AllowDouble(uc).To(ReceiveCallTo("Update").With(user).AndReturn(nil))

				Expect(c.UpdateUser(user)).Should(Succeed())
			})
			It("regular", func() {
				user := model.UserForUpdate{
					Username: "root",
					Password: "1234",
					Info: model.UserInfoForUpdate{
						Firstname: "m",
						Lastname:  "m",
						Surname:   "mm",
						Group:     "p-14",
					},
				}
				AllowDouble(uc).To(ReceiveCallTo("Update").With(user).AndReturn(errors.New("some error")))

				Expect(c.UpdateUser(user)).ShouldNot(Succeed())
			})
		})

		Context("testing UpdateBalance", func() {
			It("regular", func() {
				AllowDouble(uc).To(ReceiveCallTo("UpdateBalance").With("root", 600).AndReturn(nil))

				Expect(c.UpdateBalance("root", 600)).Should(Succeed())
			})
			It("error", func() {
				AllowDouble(uc).To(ReceiveCallTo("UpdateBalance").With("root", 600).AndReturn(errors.New("some error")))

				Expect(c.UpdateBalance("root", 600)).ShouldNot(Succeed())
			})
		})
	})
})
