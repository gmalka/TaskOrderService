package mygrpc_test

import (
	"context"
	"log"
	"net"
	"userService/build/proto"
	"userService/pkg/model"

	mygrpc "userService/transport/grpc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var _ = Describe("Mygrpc", func() {
	var cli mygrpc.RemoteOrderClient

	BeforeEach(func() {
		lis := bufconn.Listen(1024 * 1024)
		srv := grpc.NewServer()
		t := &TestGrpcServerDouble{}
		proto.RegisterTaskOrderServiceServer(srv, t)

		go func() {
			if err := srv.Serve(lis); err != nil {
				log.Fatalf("srv.Serve %v", err)
			}
		}()

		dialer := func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}
		conn, _ := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
		cli, _ = mygrpc.NewGrpcClient(conn)
	})

	Context("testing GRPC", func() {
		Context("test GetAllTasksWithoutAnswers", func() {
			It("regular", func() {
				Expect(cli.GetAllTasksWithoutAnswers(1)).To(Equal([]model.TaskWithoutAnswer{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 15},
				}))
			})
			It("error", func() {
				_, err := cli.GetAllTasksWithoutAnswers(2)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("test GetOrdersForUser", func() {
			It("regular", func() {
				Expect(cli.GetOrdersForUser("root", 1)).To(Equal([]model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 0},
				}))
			})
			It("error", func() {
				_, err := cli.GetOrdersForUser("root", 2)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("test GetAllTasks", func() {
			It("regular", func() {
				Expect(cli.GetAllTasks()).To(Equal([]model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2},
				}))
			})
		})

		Context("test CheckAndGetTask", func() {
			It("regular", func() {
				Expect(cli.CheckAndGetTask("root", 1)).To(Equal(model.TaskOrderInfo{Id: 1, Price: 500, Answer: 1}))
			})
			It("error", func() {
				_, err := cli.CheckAndGetTask("root", 2)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("test BuyTaskAnswer", func() {
			It("regular", func() {
				Expect(cli.BuyTaskAnswer("root", 1)).Should(Succeed())
			})
			It("error", func() {
				Expect(cli.BuyTaskAnswer("root", 2)).ShouldNot(Succeed())
			})
		})

		Context("test CreateNewTask", func() {
			It("regular", func() {
				Expect(cli.CreateNewTask(model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 400, Answer: 2})).Should(Succeed())
			})
			It("error", func() {
				Expect(cli.CreateNewTask(model.Task{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 400, Answer: 2})).ShouldNot(Succeed())
			})
		})

		Context("test UpdatePriceOfTask", func() {
			It("regular", func() {
				Expect(cli.UpdatePriceOfTask(1, 200)).Should(Succeed())
			})
			It("error", func() {
				Expect(cli.UpdatePriceOfTask(2, 200)).ShouldNot(Succeed())
			})
		})

		Context("test DeleteOrdersForUser", func() {
			It("regular", func() {
				Expect(cli.DeleteOrdersForUser("root")).Should(Succeed())
			})
			It("error", func() {
				Expect(cli.DeleteOrdersForUser("gmalka")).ShouldNot(Succeed())
			})
		})

		Context("test DeleteTask", func() {
			It("regular", func() {
				Expect(cli.DeleteTask(1)).Should(Succeed())
			})
			It("error", func() {
				Expect(cli.DeleteTask(2)).ShouldNot(Succeed())
			})
		})
	})
})
