package mygrpc

import (
	"TaskService/build/proto"
	"TaskService/pkg/database"
	"TaskService/pkg/model"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type RemoteOrderServer interface {
// 	Ping(ctx context.Context, req *proto.None) (*proto.None, error)
// 	getOrdersForUser(req *proto.UserOrders, stream proto.TaskOrderService_GetOrdersForUserServer) error
// 	getAllTasks(req *proto.None, stream proto.TaskOrderService_GetOrdersForUserServer) error
// 	getTask(req *proto.OrderTask) (proto.TaskOrderInfo, error)
// 	buyTaskAnswer(req *proto.UserBuyAnswer) (proto.None, error)
// 	createNewTask(req *proto.Task) (proto.None, error)
// 	updatePriceOfTask(req *proto.TaskForUpdate) (proto.None, error)
// 	deleteOrdersForUser(req *proto.UserId) (proto.None, error)
// 	deleteTask(req *proto.OrderTask) (proto.None, error)

// 	mustEmbedUnimplementedTaskOrderServiceServer()
// }

type Log struct {
	Err *log.Logger
	Inf *log.Logger
}

type grpcServ struct {
	bd     database.DatabaseService
	logger Log
	proto.UnimplementedTaskOrderServiceServer
}

func NewGrpcServer(bd database.DatabaseService, logger Log) *grpc.Server {
	grpcServer := grpc.NewServer()
	proto.RegisterTaskOrderServiceServer(grpcServer, &grpcServ{bd: bd, logger: logger})
	return grpcServer
}

func (g grpcServ) Ping(ctx context.Context, req *proto.None) (*proto.None, error) {
	return &proto.None{}, nil
}

func (g grpcServ) GetAllTasksWithoutAnswers(req *proto.Page, stream proto.TaskOrderService_GetAllTasksWithoutAnswersServer) error {
	data, err := g.bd.GetAllTasksWithoutAnswers(int(req.Page))
	if err != nil {
		g.logger.Err.Printf("can't get all tasks without answers: %v", err)
		return err
	}

	for _, s := range data {
		select {
		case <-stream.Context().Done():
			g.logger.Err.Printf("get all stream has ended: %v", err)
			return status.Error(codes.Canceled, "stream has ended")
		default:
			err := stream.SendMsg(&proto.Task{
				Id:     int64(s.Id),
				Count:  int64(s.Count),
				Height: s.Heights,
				Price:  int64(s.Price),
			})
			if err != nil {
				g.logger.Err.Printf("can't send message: %v", err)
				return status.Error(codes.Canceled, err.Error())
			}
		}
	}
	return nil
}

func (g grpcServ) GetOrdersForUser(req *proto.UserOrders, stream proto.TaskOrderService_GetOrdersForUserServer) error {
	data, err := g.bd.GetAllTasksOfUser(req.Username, int(req.Page))
	if err != nil {
		g.logger.Err.Printf("can't get all tasks of user %s: %v", req.Username, err)
		return err
	}

	for _, s := range data {
		select {
		case <-stream.Context().Done():
			g.logger.Err.Printf("get all stream has ended: %v", err)
			return status.Error(codes.Canceled, "stream has ended")
		default:
			err := stream.SendMsg(&proto.Task{
				Id:     int64(s.Id),
				Count:  int64(s.Count),
				Height: s.Heights,
				Price:  int64(s.Price),
				Answer: int64(s.Answer),
			})
			if err != nil {
				g.logger.Err.Printf("can't send message: %v", err)
				return status.Error(codes.Canceled, err.Error())
			}
		}
	}
	return nil
}

func (g grpcServ) GetAllTasks(req *proto.None, stream proto.TaskOrderService_GetAllTasksServer) error {
	data, err := g.bd.GetAllTasks()
	if err != nil {
		g.logger.Err.Printf("can't get all tasks: %v", err)
		return err
	}

	for _, s := range data {
		select {
		case <-stream.Context().Done():
			g.logger.Err.Printf("get all stream has ended: %v", err)
			return status.Error(codes.Canceled, "stream has ended")
		default:
			err := stream.SendMsg(&proto.Task{
				Id:     int64(s.Id),
				Count:  int64(s.Count),
				Height: s.Heights,
				Price:  int64(s.Price),
				Answer: int64(s.Answer),
			})
			if err != nil {
				g.logger.Err.Printf("can't send message: %v", err)
				return status.Error(codes.Canceled, err.Error())
			}
		}
	}
	return nil
}

func (g grpcServ) CheckAndGetTask(ctx context.Context, req *proto.UsernameAndId) (*proto.TaskOrderInfo, error) {
	task, err := g.bd.CheckAndGetTask(req.Username, int(req.Id))

	if err != nil {
		g.logger.Err.Printf("can't check and get task %d by user %s: %v", req.Id, req.Username, err)
	}

	return &proto.TaskOrderInfo{
		Answer: int64(task.Answer),
		Price:  int64(task.Price),
	}, err
}

func (g grpcServ) BuyTaskAnswer(ctx context.Context, req *proto.UsernameAndId) (*proto.None, error) {
	err := g.bd.BuyTaskAnswer(model.UsersPurchase{
		Username: req.Username,
		OrderId:  int(req.Id),
	})

	if err != nil {
		g.logger.Err.Printf("can't buy task answer %d by user %s: %v", req.Id, req.Username, err)
	}

	return &proto.None{}, err
}

func (g grpcServ) CreateNewTask(ctx context.Context, req *proto.Task) (*proto.None, error) {
	err := g.bd.CreateTask(model.Task{
		Id:      int(req.Id),
		Count:   int(req.Count),
		Heights: req.Height,
		Price:   int(req.Price),
		Answer:  int(req.Answer),
	})

	if err != nil {
		g.logger.Err.Printf("can't create task: %v", err)
	}

	return &proto.None{}, err
}

func (g grpcServ) UpdatePriceOfTask(ctx context.Context, req *proto.TaskForUpdate) (*proto.None, error) {
	err := g.bd.ChangeTaskPrice(int(req.Id), int(req.Price))
	if err != nil {
		g.logger.Err.Printf("can't update price of task %d: %v", req.Id, err)
	}

	return &proto.None{}, err
}

func (g grpcServ) DeleteOrdersForUser(ctx context.Context, req *proto.UserId) (*proto.None, error) {
	err := g.bd.DeleteAllTasksOfUser(req.Username)
	if err != nil {
		g.logger.Err.Printf("can't delete orders for user %s: %v", req.Username, err)
	}

	return &proto.None{}, err
}

func (g grpcServ) DeleteTask(ctx context.Context, req *proto.OrderTask) (*proto.None, error) {
	err := g.bd.DeleteTask(int(req.Id))
	if err != nil {
		g.logger.Err.Printf("can't delete task %d: %v", req.Id, err)
	}

	return &proto.None{}, err
}
