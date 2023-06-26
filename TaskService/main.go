package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"taskServer/database/postgres"
	postgresservice "taskServer/database/postgres/postgres_service"
	ordercontroller "taskServer/order_controller"
	mygrpc "taskServer/transport/grpc"
	"taskServer/transport/rest"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	config := postgres.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Sslmode:  os.Getenv("DB_SSLMODE"),
	}

	loggerErr := log.New(os.Stderr, "ERROR:\t ", log.Lshortfile|log.Ltime)
	loggerInfo := log.New(os.Stdout, "INFO:\t ", log.Lshortfile|log.Ltime)

	log := rest.Log{
		Err: loggerErr,
		Inf: loggerInfo,
	}

	db, err := postgres.NewPostgresConnection(config)
	if err != nil {
		log.Err.Println(err)
		return
	}
	defer db.Close()

	service := postgresservice.NewPostgresService(db)
	ordercontroller := ordercontroller.NewUserController(service)
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("GRPC_URL"), os.Getenv("GRPC_PORT")))

	grpcServer := mygrpc.NewGrpcServer(ordercontroller, mygrpc.Log(log))

	h := rest.NewHandler(ordercontroller, log)

	RunServer(fmt.Sprintf("%s:%s", os.Getenv("REST_URL"), os.Getenv("REST_PORT")),
		list, h.InitRouter(), grpcServer, log)
}

func RunServer(addr string, list net.Listener, h http.Handler, grpsServ *grpc.Server, log rest.Log) {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	log.Inf.Println("waiting for connections...")
	go srv.ListenAndServe()
	go grpsServ.Serve(list)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Inf.Println("Shutdown Server...")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	srv.Shutdown(ctx)
	grpsServ.Stop()
	log.Inf.Println("Server exited")
}
