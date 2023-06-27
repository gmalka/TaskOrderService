package mygrpc_test

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"taskServer/build/proto"
	"taskServer/model"
	mygrpc "taskServer/transport/grpc"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener

func getDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

var _ = Describe("GrpcServ", func() {
	var client proto.TaskOrderServiceClient
	var conn *grpc.ClientConn
	var err error
	var controller ControllerDouble

	BeforeEach(func() {
		loggerErr := log.New(ioutil.Discard, "ERROR:\t ", log.Lshortfile|log.Ltime)
		loggerInfo := log.New(ioutil.Discard, "INFO:\t ", log.Lshortfile|log.Ltime)
		logger := mygrpc.Log{loggerErr, loggerInfo}

		lis = bufconn.Listen(1024 * 1024)

		controller = NewDbDouble()
		s := mygrpc.NewGrpcServer(controller, logger)

		go func() {
			if err := s.Serve(lis); err != nil {
				log.Fatalf("Server exited with error: %v", err)
			}
		}()
		conn, err = grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getDialer), grpc.WithInsecure())
		Expect(err).Should(Succeed())
		client = proto.NewTaskOrderServiceClient(conn)
	})

	Context("testing grpc", func() {
		Context("testing Ping", func() {
			It("regular", func() {
				_, err := client.Ping(context.Background(), &proto.None{})
				Expect(err).Should(Succeed())
			})
		})

		Context("testing GetAllTasks", func() {
			It("regular", func() {
				wanted := []model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2},
				}
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasks").With().AndReturn(wanted, nil))

				req, err := client.GetAllTasks(context.Background(), &proto.None{})
				Expect(err).Should(Succeed())

				result := make([]model.Task, 0, 10)
				for {
					resp, err := req.Recv()
					if err == io.EOF {
						break
					}

					Expect(err).Should(Succeed())

					result = append(result, model.Task{
						Id:      int(resp.Id),
						Count:   int(resp.Count),
						Heights: resp.Height,
						Price:   int(resp.Price),
						Answer:  int(resp.Answer),
					})
				}

				Expect(result).To(Equal(wanted))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasks").With().AndReturn(nil, errors.New("some text")))

				req, err := client.GetAllTasks(context.Background(), &proto.None{})
				Expect(err).Should(Succeed())

				for {
					_, err := req.Recv()
					if err == io.EOF {
						break
					}

					Expect(err).ShouldNot(Succeed())
					break
				}
			})
		})

		Context("testing GetAllTasksWithoutAnswers", func() {
			It("regular", func() {
				wanted := []model.TaskWithoutAnswer{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500},
					{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 500},
				}
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasksWithoutAnswers").With(1).AndReturn(wanted, nil))
				req, err := client.GetAllTasksWithoutAnswers(context.Background(), &proto.Page{Page: 1})

				Expect(err).Should(Succeed())

				result := make([]model.TaskWithoutAnswer, 0, 10)
				for {
					resp, err := req.Recv()
					if err == io.EOF {
						break
					}

					Expect(err).Should(Succeed())

					result = append(result, model.TaskWithoutAnswer{
						Id:      int(resp.Id),
						Count:   int(resp.Count),
						Heights: resp.Height,
						Price:   int(resp.Price),
					})
				}

				Expect(result).To(Equal(wanted))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasksWithoutAnswers").With(1).AndReturn(nil, errors.New("some error")))
				req, err := client.GetAllTasksWithoutAnswers(context.Background(), &proto.Page{Page: 1})

				Expect(err).Should(Succeed())

				for {
					_, err := req.Recv()
					if err == io.EOF {
						break
					}

					Expect(err).ShouldNot(Succeed())
					return
				}
			})
		})

		Context("testing GetOrdersForUser", func() {
			It("regular", func() {
				wanted := []model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 4},
					{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 6},
				}
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasksOfUser").With("root", 1).AndReturn(wanted, nil))
				req, err := client.GetOrdersForUser(context.Background(), &proto.UserOrders{
					Username: "root",
					Page:     1,
				})
				Expect(err).Should(Succeed())

				result := make([]model.Task, 0, 10)
				for {
					resp, err := req.Recv()
					if err == io.EOF {
						break
					}

					Expect(err).Should(Succeed())

					result = append(result, model.Task{
						Id:      int(resp.Id),
						Count:   int(resp.Count),
						Heights: resp.Height,
						Price:   int(resp.Price),
						Answer:  int(resp.Answer),
					})
				}

				Expect(result).To(Equal(wanted))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasksOfUser").With("root", 1).AndReturn(nil, errors.New("some error")))
				req, err := client.GetOrdersForUser(context.Background(), &proto.UserOrders{
					Username: "root",
					Page:     1,
				})
				Expect(err).Should(Succeed())

				for {
					_, err := req.Recv()
					if err == io.EOF {
						break
					}

					Expect(err).ShouldNot(Succeed())
					break
				}

			})
		})

		Context("testing CheckAndGetTask", func() {
			It("regular", func() {
				in := model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 4}
				AllowDouble(controller).To(ReceiveCallTo("CheckAndGetTask").With("root", 1).AndReturn(in, nil))
				req, err := client.CheckAndGetTask(context.Background(), &proto.UsernameAndId{
					Username: "root",
					Id:       1,
				})
				Expect(err).Should(Succeed())

				Expect(req.Answer).To(Equal(int64(4)))
				Expect(req.Price).To(Equal(int64(500)))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("CheckAndGetTask").With("root", 1).AndReturn(nil, errors.New("some error")))
				_, err := client.CheckAndGetTask(context.Background(), &proto.UsernameAndId{
					Username: "root",
					Id:       1,
				})
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing BuyTaskAnswer", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("BuyTaskAnswer").With(model.UsersPurchase{Username: "root", OrderId: 1}).AndReturn(nil))
				_, err := client.BuyTaskAnswer(context.Background(), &proto.UsernameAndId{
					Username: "root",
					Id:       1,
				})
				Expect(err).Should(Succeed())
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("BuyTaskAnswer").With(model.UsersPurchase{Username: "root", OrderId: 1}).AndReturn(errors.New("some error")))
				_, err := client.BuyTaskAnswer(context.Background(), &proto.UsernameAndId{
					Username: "root",
					Id:       1,
				})
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing CreateNewTask", func() {
			It("regular", func() {
				in := model.Task{
					Id:      0,
					Count:   4,
					Heights: []int64{1, 2, 3, 4},
					Price:   500,
					Answer:  2,
				}
				AllowDouble(controller).To(ReceiveCallTo("CreateTask").With(in).AndReturn(nil))
				_, err := client.CreateNewTask(context.Background(), &proto.Task{
					Id:     int64(in.Id),
					Count:  int64(in.Count),
					Height: in.Heights,
					Price:  int64(in.Price),
					Answer: int64(in.Answer),
				})

				Expect(err).Should(Succeed())
			})
			It("error", func() {
				in := model.Task{
					Id:      0,
					Count:   4,
					Heights: []int64{1, 2, 3, 4},
					Price:   500,
					Answer:  2,
				}
				AllowDouble(controller).To(ReceiveCallTo("CreateTask").With(in).AndReturn(errors.New("some error")))
				_, err := client.CreateNewTask(context.Background(), &proto.Task{
					Id:     int64(in.Id),
					Count:  int64(in.Count),
					Height: in.Heights,
					Price:  int64(in.Price),
					Answer: int64(in.Answer),
				})

				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing UpdatePriceOfTask", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("ChangeTaskPrice").With(1, 500).AndReturn(nil))
				_, err := client.UpdatePriceOfTask(context.Background(), &proto.TaskForUpdate{
					Id:    1,
					Price: 500,
				})

				Expect(err).Should(Succeed())
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("ChangeTaskPrice").With(1, 500).AndReturn(errors.New("some error")))
				_, err := client.UpdatePriceOfTask(context.Background(), &proto.TaskForUpdate{
					Id:    1,
					Price: 500,
				})

				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing DeleteOrdersForUser", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteAllTasksOfUser").With("root").AndReturn(nil))
				_, err := client.DeleteOrdersForUser(context.Background(), &proto.UserId{
					Username: "root",
				})

				Expect(err).Should(Succeed())
			})

			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteAllTasksOfUser").With("root").AndReturn(errors.New("some error")))
				_, err := client.DeleteOrdersForUser(context.Background(), &proto.UserId{
					Username: "root",
				})

				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing DeleteTask", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteTask").With(1).AndReturn(nil))
				_, err := client.DeleteTask(context.Background(), &proto.OrderTask{
					Id: 1,
				})

				Expect(err).Should(Succeed())
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteTask").With(1).AndReturn(errors.New("some error")))
				_, err := client.DeleteTask(context.Background(), &proto.OrderTask{
					Id: 1,
				})

				Expect(err).ShouldNot(Succeed())
			})
		})
	})

	AfterEach(func() {
		defer conn.Close()
	})
})
