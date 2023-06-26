package mygrpc_test

import (
	"context"
	"errors"
	"testing"
	"userService/build/proto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGrpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grpc Suite")
}

type TestGrpcServerDouble struct {
	proto.UnimplementedTaskOrderServiceServer
}

func (t TestGrpcServerDouble) Ping(ctx context.Context, req *proto.None) (*proto.None, error) {
	return &proto.None{}, nil
}

func (t TestGrpcServerDouble) GetAllTasksWithoutAnswers(req *proto.Page, stream proto.TaskOrderService_GetAllTasksWithoutAnswersServer) error {
	switch req.Page {
	case 1:
		stream.SendMsg(&proto.Task{
			Id:     int64(1),
			Count:  int64(2),
			Height: []int64{1, 2},
			Price:  int64(15),
		})
	case 2:
		return errors.New("some error")
	}

	return nil
}

func (t TestGrpcServerDouble) GetOrdersForUser(req *proto.UserOrders, stream proto.TaskOrderService_GetOrdersForUserServer) error {
	switch req.Page {
	case 1:
		stream.SendMsg(&proto.Task{
			Id:     int64(1),
			Count:  int64(2),
			Height: []int64{1, 2},
			Price:  int64(500),
			Answer: int64(3),
		})
	case 2:
		return errors.New("some error")
	}

	return nil
}

func (t TestGrpcServerDouble) GetAllTasks(req *proto.None, stream proto.TaskOrderService_GetAllTasksServer) error {
	stream.SendMsg(&proto.Task{
		Id:     int64(1),
		Count:  int64(2),
		Height: []int64{1, 2},
		Price:  int64(500),
		Answer: int64(2),
	})

	return nil
}

func (t TestGrpcServerDouble) CheckAndGetTask(ctx context.Context, req *proto.UsernameAndId) (*proto.TaskOrderInfo, error) {
	switch req.Id {
	case 1:
		return &proto.TaskOrderInfo{
			Answer: int64(1),
			Price:  int64(500),
		}, nil
	case 2:
		return nil, errors.New("some error")
	}

	return nil, nil
}

func (t TestGrpcServerDouble) BuyTaskAnswer(ctx context.Context, req *proto.UsernameAndId) (*proto.None, error) {
	switch req.Id {
	case 1:
		return &proto.None{}, nil
	case 2:
		return &proto.None{}, errors.New("some error")
	}

	return nil, nil
}

func (t TestGrpcServerDouble) CreateNewTask(ctx context.Context, req *proto.Task) (*proto.None, error) {
	switch req.Id {
	case 1:
		return &proto.None{}, nil
	case 2:
		return &proto.None{}, errors.New("some error")
	}

	return nil, nil
}

func (t TestGrpcServerDouble) UpdatePriceOfTask(ctx context.Context, req *proto.TaskForUpdate) (*proto.None, error) {
	switch req.Id {
	case 1:
		return &proto.None{}, nil
	case 2:
		return &proto.None{}, errors.New("some error")
	}

	return nil, nil
}

func (t TestGrpcServerDouble) DeleteOrdersForUser(ctx context.Context, req *proto.UserId) (*proto.None, error) {
	switch req.Username {
	case "root":
		return &proto.None{}, nil
	case "gmalka":
		return &proto.None{}, errors.New("some error")
	}

	return nil, nil
}

func (t TestGrpcServerDouble) DeleteTask(ctx context.Context, req *proto.OrderTask) (*proto.None, error) {
	switch req.Id {
	case 1:
		return &proto.None{}, nil
	case 2:
		return &proto.None{}, errors.New("some error")
	}

	return nil, nil
}
