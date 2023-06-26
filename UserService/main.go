package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"userService/internal/auth"
	"userService/internal/database/postgres"
	postgresservice "userService/internal/database/postgres/postgres_service"
	mygrpc "userService/internal/transport/grpc"
	"userService/internal/transport/rest"
	usercontroller "userService/internal/user_controller"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	tokenManager := auth.NewAuthService(os.Getenv("ACCESS_SECRET"), os.Getenv("REFRESH_SECRET"))

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

	userDB := postgresservice.NewPostgresService(db)
	userController := usercontroller.NewUserController(userDB)
	grpcController, err := mygrpc.NewGrpcClient(os.Getenv("GRPC_URL"), os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Err.Println(err)
		return
	}

	handler := rest.NewHandler(userController, tokenManager, grpcController, auth.NewPasswordManager(), log)

	RunServer(fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT")), handler.InitRouter(), log)
}

func RunServer(addr string, h http.Handler, log rest.Log) {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	log.Inf.Println("waiting for connections...")
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Inf.Println("Shutdown Server...")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	srv.Shutdown(ctx)
	log.Inf.Println("Server exited")
}

/*
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative build/proto/order_service.proto

swagger generate spec -o public/swagger.json --scan-models
*/
