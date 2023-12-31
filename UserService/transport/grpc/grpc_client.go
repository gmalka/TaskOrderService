package mygrpc

import (
	"context"
	"fmt"
	"io"
	"time"
	"userService/build/proto"
	"userService/pkg/model"

	"google.golang.org/grpc"
)

type RemoteOrderClient interface {
	UpdatePriceOfTask(id, price int) error
	CreateNewTask(task model.TaskWithoutAnswer) error
	CheckAndGetTask(username string, id int) (model.TaskOrderInfo, error)
	BuyTaskAnswer(username string, taskId int) error
	GetAllTasks() ([]model.Task, error)
	GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error)
	GetOrdersForUser(username string, page int) ([]model.Task, error)
	DeleteOrdersForUser(username string) error
	DeleteTask(taskId int) error
}

type grpcClient struct {
	client proto.TaskOrderServiceClient
}

func NewGrpcClient(conn *grpc.ClientConn) (RemoteOrderClient, error) {
	var err error

	client := proto.NewTaskOrderServiceClient(conn)

	for i := 0; i < 5; i++ {
		_, err = client.Ping(context.Background(), &proto.None{})
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("can't ping to grpc server: %v", err)
	}

	return grpcClient{client: client}, nil
}

func (g grpcClient) GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error) {
	res, err := g.client.GetAllTasksWithoutAnswers(context.Background(), &proto.Page{
		Page: int64(page),
	})
	if err != nil {
		return nil, fmt.Errorf("can't get all tasks: %v", err)
	}

	result := make([]model.TaskWithoutAnswer, 0, 10)
	for {
		resp, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("can't get task: %v", err)
		}

		result = append(result, model.TaskWithoutAnswer{
			Id:      int(resp.Id),
			Count:   int(resp.Count),
			Heights: resp.Height,
			Price:   int(resp.Price),
		})
	}

	return result, nil
}

func (g grpcClient) DeleteTask(taskId int) error {
	_, err := g.client.DeleteTask(context.Background(), &proto.OrderTask{
		Id: int64(taskId),
	})
	if err != nil {
		return fmt.Errorf("can't delete task %d: %v", taskId, err)
	}

	return nil
}

func (g grpcClient) DeleteOrdersForUser(username string) error {
	_, err := g.client.DeleteOrdersForUser(context.Background(), &proto.UserId{
		Username: username,
	})
	if err != nil {
		return fmt.Errorf("can't delete orders for user %s: %v", username, err)
	}

	return nil
}

func (g grpcClient) BuyTaskAnswer(username string, taskId int) error {
	_, err := g.client.BuyTaskAnswer(context.Background(), &proto.UsernameAndId{
		Username: username,
		Id:       int64(taskId),
	})
	if err != nil {
		return fmt.Errorf("can't buy tasks answer for user %s: %v", username, err)
	}

	return nil
}

func (g grpcClient) UpdatePriceOfTask(id, price int) error {
	_, err := g.client.UpdatePriceOfTask(context.Background(), &proto.TaskForUpdate{
		Id:    int64(id),
		Price: int64(price),
	})
	if err != nil {
		return fmt.Errorf("can't update tasks price by id %d: %v", id, err)
	}

	return nil
}

func (g grpcClient) CreateNewTask(task model.TaskWithoutAnswer) error {
	_, err := g.client.CreateNewTask(context.Background(), &proto.TaskWithoutAnswer{
		Id:     int64(task.Id),
		Count:  int64(task.Count),
		Height: task.Heights,
		Price:  int64(task.Price),
	})
	if err != nil {
		return fmt.Errorf("can't create new task: %v", err)
	}

	return nil
}

func (g grpcClient) CheckAndGetTask(username string, id int) (model.TaskOrderInfo, error) {
	var task model.TaskOrderInfo

	res, err := g.client.CheckAndGetTask(context.Background(), &proto.UsernameAndId{
		Username: username,
		Id:       int64(id),
	})
	if err != nil {
		return task, fmt.Errorf("can't get task %d: %v", id, err)
	}

	task.Id = id
	task.Price = int(res.Price)
	task.Answer = int(res.Answer)

	return task, nil
}

func (g grpcClient) GetAllTasks() ([]model.Task, error) {
	res, err := g.client.GetAllTasks(context.Background(), &proto.None{})
	if err != nil {
		return nil, fmt.Errorf("can't get all tasks: %v", err)
	}

	result := make([]model.Task, 0, 10)
	for {
		resp, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("can't get task: %v", err)
		}

		result = append(result, model.Task{
			Id:      int(resp.Id),
			Count:   int(resp.Count),
			Heights: resp.Height,
			Price:   int(resp.Price),
			Answer:  int(resp.Answer),
		})
	}

	return result, nil
}

func (g grpcClient) GetOrdersForUser(username string, page int) ([]model.Task, error) {
	res, err := g.client.GetOrdersForUser(context.Background(), &proto.UserOrders{
		Username: username,
		Page:     int64(page),
	})
	if err != nil {
		return nil, fmt.Errorf("can't get order for user %s: %v", username, err)
	}

	result := make([]model.Task, 0, 10)
	for {
		resp, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("can't get users task: %v", err)
		}

		result = append(result, model.Task{
			Id:      int(resp.Id),
			Count:   int(resp.Count),
			Heights: resp.Height,
			Price:   int(resp.Price),
		})
	}

	return result, nil
}
