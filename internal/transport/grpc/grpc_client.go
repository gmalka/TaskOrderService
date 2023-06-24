package grpc

import (
	"context"
	"fmt"
	"io"
	"userService/build/proto/build/proto"
	"userService/internal/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RemoteOrderClient interface {
	UpdatePriceOfTask(id, price int) error
	CreateNewTask(task model.Task) error
	GetTask(id int) (model.TaskOrderInfo, error)
	BuyTaskAnswer(username string, taskId int) error
	GetAllTasks() ([]model.Task, error)
	GetOrdersForUser(username string, page int) ([]model.Task, error)
	DeleteOrdersForUser(username string) error
	DeleteTask(taskId int) error
}

type grpcClient struct {
	client proto.TaskOrderServiceClient
}

func NewGrpcClient(ip, port string) (RemoteOrderClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	path := fmt.Sprintf("%s:%s", ip, port)
	conn, err := grpc.Dial(path, opts...)
	if err != nil {
		return nil, fmt.Errorf("can't connect by grpc to path %s: %v", path, err)
	}

	client := proto.NewTaskOrderServiceClient(conn)

	/*_, err = client.Ping(context.Background(), &proto.None{})
	if err != nil {
		return nil, fmt.Errorf("can't ping to grpc server by path %s: %v", path, err)
	}*/

	return grpcClient{client: client}, nil
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
	_, err := g.client.BuyTaskAnswer(context.Background(), &proto.UserBuyAnswer{
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

func (g grpcClient) CreateNewTask(task model.Task) error {
	_, err := g.client.CreateNewTask(context.Background(), &proto.Task{
		Id:     int64(task.Id),
		Count:  int64(task.Count),
		Height: task.Heights,
		Price:  int64(task.Price),
		Answer: int64(task.Answer),
	})
	if err != nil {
		return fmt.Errorf("can't create new task: %v", err)
	}

	return nil
}

func (g grpcClient) GetTask(id int) (model.TaskOrderInfo, error) {
	var task model.TaskOrderInfo

	res, err := g.client.GetTask(context.Background(), &proto.OrderTask{
		Id: int64(id),
	})
	if err != nil {
		return task, fmt.Errorf("can't get task by id %d: %v", id, err)
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
