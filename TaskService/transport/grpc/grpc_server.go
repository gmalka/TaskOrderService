package mygrpc

import (
	"context"
	"taskServer/build/proto"
	"taskServer/database"
	"taskServer/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type RemoteOrderServer interface {
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

type grpcServ struct {
	bd database.DatabaseService
	proto.UnimplementedTaskOrderServiceServer
}

func NewGrpcServer(bd database.DatabaseService) *grpc.Server {
	grpcServer := grpc.NewServer()
	proto.RegisterTaskOrderServiceServer(grpcServer, &grpcServ{bd: bd})
	return grpcServer
}

func (g grpcServ) Ping(ctx context.Context, req *proto.None) (*proto.None, error) {
	return &proto.None{}, nil
}

func (g grpcServ) GetOrdersForUser(req *proto.UserOrders, stream proto.TaskOrderService_GetOrdersForUserServer) error {
	data, err := g.bd.GetAllTasksOfUser(req.Username, int(req.Page))
	if err != nil {
		return err
	}

	for _, s := range data {
		select {
		case <-stream.Context().Done():
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
				return status.Error(codes.Canceled, err.Error())
			}
		}
	}
	return nil
}

func (g grpcServ) GetAllTasks(req *proto.None, stream proto.TaskOrderService_GetAllTasksServer) error {
	data, err := g.bd.GetAllTasks()
	if err != nil {
		return err
	}

	for _, s := range data {
		select {
		case <-stream.Context().Done():
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
				return status.Error(codes.Canceled, err.Error())
			}
		}
	}
	return nil
}

func (g grpcServ) GetTask(ctx context.Context, req *proto.OrderTask) (*proto.TaskOrderInfo, error) {
	task, err := g.bd.GetTask(int(req.Id))

	return &proto.TaskOrderInfo{
		Answer: int64(task.Answer),
		Price:  int64(task.Price),
	}, err
}

func (g grpcServ) BuyTaskAnswer(ctx context.Context, req *proto.UserBuyAnswer) (*proto.None, error) {
	err := g.bd.BuyTaskAnswer(model.UsersPurchase{
		Username: req.Username,
		OrderId:  int(req.Id),
	})

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

	return &proto.None{}, err
}

func (g grpcServ) UpdatePriceOfTask(ctx context.Context, req *proto.TaskForUpdate) (*proto.None, error) {
	err := g.bd.ChangeTaskPrice(int(req.Id), int(req.Price))

	return &proto.None{}, err
}

func (g grpcServ) DeleteOrdersForUser(ctx context.Context, req *proto.UserId) (*proto.None, error) {
	err := g.bd.DeleteAllTasksOfUser(req.Username)

	return &proto.None{}, err
}

func (g grpcServ) DeleteTask(ctx context.Context, req *proto.OrderTask) (*proto.None, error) {
	err := g.bd.DeleteTask(int(req.Id))

	return &proto.None{}, err
}
