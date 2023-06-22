package grpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"userService/build/proto/build/proto"
	"userService/internal/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RemoteOrderService interface {
	UpdatePriceOfTask(id, price int) error
	CreateNewTask(task model.Task) error
	GetTask(id int) (model.TaskOrderInfo, error)
	BuyTaskAnswer(username string, taskId int) error
	GetAllTasks() ([]model.Task, error)
	GetOrdersForUser(username string, page int) ([]model.Task, error)
}

type grpcClient struct {
	client proto.TaskOrderServiceClient
}

func NewGrpcClient(ip, port string) (RemoteOrderService, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	path := fmt.Sprintf("%s:%s", ip, port)
	conn, err := grpc.Dial(path, opts...)
	if err != nil {
		return nil, err
	}
	
	client := proto.NewTaskOrderServiceClient(conn)
	return grpcClient{client: client}, nil
}

func (g grpcClient) BuyTaskAnswer(username string, taskId int) error {
	_, err := g.client.BuyTaskAnswer(context.Background(), &proto.UserBuyAnswer{
		Username: username,
		Id: int64(taskId),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g grpcClient) UpdatePriceOfTask(id, price int) error {
	_, err := g.client.UpdatePriceOfTask(context.Background(), &proto.TaskForUpdate{
		Id: int64(id),
		Price: int64(price),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g grpcClient) CreateNewTask(task model.Task) error {
	_, err := g.client.CreateNewTask(context.Background(), &proto.Task{
		Id: int64(task.Id),
		Count: int64(task.Count),
		Heiaghts: task.Heights,
		Price: int64(task.Price),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g grpcClient) GetTask(id int) (model.TaskOrderInfo, error) {
	var task model.TaskOrderInfo

	res, err := g.client.GetTask(context.Background(), &proto.OrderTask{
		Id: int64(id),
	})
	if err != nil {
		return task, err
	}

	task.Id = id
	task.Price = int(res.Price)
	task.Answer = int(res.Answer)

	return task, nil
}

func (g grpcClient) GetAllTasks() ([]model.Task, error) {
	res, err := g.client.GetAllTasks(context.Background(), &proto.None{})
	if err != nil {
		return nil, err
	}

	result := make([]model.Task, 0, 10)
	for {
		resp, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, errors.New("error while geting orders for user")
		}

		result = append(result, model.Task{
			Id: int(resp.Id),
			Count: int(resp.Count),
			Heights: resp.Heiaghts,
			Price: int(resp.Price),
		})
	}

	return result, nil
}

func (g grpcClient) GetOrdersForUser(username string, page int) ([]model.Task, error) {
	res, err := g.client.GetOrdersForUser(context.Background(), &proto.UserOrders{
		Username: username,
		Page: int64(page),
	})
	if err != nil {
		return nil, err
	}

	result := make([]model.Task, 0, 10)
	for {
		resp, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, errors.New("error while geting orders for user")
		}

		result = append(result, model.Task{
			Id: int(resp.Id),
			Count: int(resp.Count),
			Heights: resp.Heiaghts,
			Price: int(resp.Price),
		})
	}

	return result, nil
}