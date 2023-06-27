package ordercontroller_test

import (
	"errors"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"taskServer/model"
	ordercontroller "taskServer/order_controller"
)

var _ = Describe("OrderController", func() {
	var (
		db DbDouble
		c  ordercontroller.Controller
	)

	BeforeEach(func() {
		db = NewDbDouble()
		c = ordercontroller.NewUserController(db)
	})

	Context("order controller tests", func() {
		Context("testing GetAllTasks", func() {
			It("regular", func() {
				want := []model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15},
					{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15},
				}
				AllowDouble(db).To(ReceiveCallTo("GetAllTasks").With().AndReturn(want, nil))

				Expect(c.GetAllTasks()).To(Equal(want))
			})
			It("error", func() {
				AllowDouble(db).To(ReceiveCallTo("GetAllTasks").With().AndReturn(nil, errors.New("some error")))

				_, err := c.GetAllTasks()
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing CreateTask", func() {
			It("regular", func() {
				in := model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15}

				AllowDouble(db).To(ReceiveCallTo("CreateTask").With(in).AndReturn(nil))

				Expect(c.CreateTask(in)).Should(Succeed())
			})
			It("error", func() {
				in := model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15}
				AllowDouble(db).To(ReceiveCallTo("CreateTask").With(in).AndReturn(errors.New("some error")))

				err := c.CreateTask(in)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing CheckAndGetTask", func() {
			It("regular", func() {
				inUsername := "root"
				inId := 2
				want := model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15}

				AllowDouble(db).To(ReceiveCallTo("CheckAndGetTask").With(inUsername, inId).AndReturn(want, nil))

				Expect(c.CheckAndGetTask(inUsername, inId)).To(Equal(want))
			})
			It("error", func() {
				inUsername := "root"
				inId := 2
				want := model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15}
				AllowDouble(db).To(ReceiveCallTo("CheckAndGetTask").With(inUsername, inId).AndReturn(want, errors.New("some error")))

				_, err := c.CheckAndGetTask(inUsername, inId)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing ChangeTaskPrice", func() {
			It("regular", func() {
				inId := 2
				inPrice := 500

				AllowDouble(db).To(ReceiveCallTo("ChangeTaskPrice").With(inId, inPrice).AndReturn(nil))

				Expect(c.ChangeTaskPrice(inId, inPrice)).Should(Succeed())
			})
			It("error", func() {
				inId := 2
				inPrice := 500
				AllowDouble(db).To(ReceiveCallTo("ChangeTaskPrice").With(inId, inPrice).AndReturn(errors.New("some error")))

				err := c.ChangeTaskPrice(inId, inPrice)
				Expect(err).ShouldNot(Succeed())
			})

			Context("testing DeleteTask", func() {
				It("regular", func() {
					inId := 2

					AllowDouble(db).To(ReceiveCallTo("DeleteTask").With(inId).AndReturn(nil))

					Expect(c.DeleteTask(inId)).Should(Succeed())
				})
				It("error", func() {
					inId := 2
					AllowDouble(db).To(ReceiveCallTo("DeleteTask").With(inId).AndReturn(errors.New("some error")))

					err := c.DeleteTask(inId)
					Expect(err).ShouldNot(Succeed())
				})
			})

			Context("testing GetAllTasksOfUser", func() {
				It("regular", func() {
					inUsername := "root"
					inId := 2
					want := []model.Task{
						{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15},
						{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 15},
					}
					AllowDouble(db).To(ReceiveCallTo("GetAllTasksOfUser").With(inUsername, inId).AndReturn(want, nil))

					Expect(c.GetAllTasksOfUser(inUsername, inId)).To(Equal(want))
				})
				It("error", func() {
					inUsername := "root"
					inId := 2
					AllowDouble(db).To(ReceiveCallTo("GetAllTasksOfUser").With(inUsername, inId).AndReturn(nil, errors.New("some error")))

					_, err := c.GetAllTasksOfUser(inUsername, inId)
					Expect(err).ShouldNot(Succeed())
				})
			})

			Context("testing GetAllTasksWithoutAnswers", func() {
				It("regular", func() {
					inPage := 2
					want := []model.TaskWithoutAnswer{
						{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500},
						{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 500},
					}
					AllowDouble(db).To(ReceiveCallTo("GetAllTasksWithoutAnswers").With(inPage).AndReturn(want, nil))

					Expect(c.GetAllTasksWithoutAnswers(inPage)).To(Equal(want))
				})
				It("error", func() {
					inPage := 2
					AllowDouble(db).To(ReceiveCallTo("GetAllTasksWithoutAnswers").With(inPage).AndReturn(nil, errors.New("some error")))

					_, err := c.GetAllTasksWithoutAnswers(inPage)
					Expect(err).ShouldNot(Succeed())
				})
			})

			Context("testing DeleteAllTasksOfUser", func() {
				It("regular", func() {
					inUsername := "root"
					AllowDouble(db).To(ReceiveCallTo("DeleteAllTasksOfUser").With(inUsername).AndReturn(nil))

					Expect(c.DeleteAllTasksOfUser(inUsername)).Should(Succeed())
				})
				It("error", func() {
					inUsername := "root"
					AllowDouble(db).To(ReceiveCallTo("DeleteAllTasksOfUser").With(inUsername).AndReturn(errors.New("some error")))

					err := c.DeleteAllTasksOfUser(inUsername)
					Expect(err).ShouldNot(Succeed())
				})
			})
			
			Context("testing BuyTaskAnswer", func() {
				It("regular", func() {
					want := model.UsersPurchase{Username: "root", OrderId: 1}
					
					AllowDouble(db).To(ReceiveCallTo("BuyTaskAnswer").With(want).AndReturn(nil))

					Expect(c.BuyTaskAnswer(want)).Should(Succeed())
				})
				It("error", func() {
					want := model.UsersPurchase{Username: "root", OrderId: 1}
					AllowDouble(db).To(ReceiveCallTo("BuyTaskAnswer").With(want).AndReturn(errors.New("some error")))

					err := c.BuyTaskAnswer(want)
					Expect(err).ShouldNot(Succeed())
				})
			})
		})
	})
})
