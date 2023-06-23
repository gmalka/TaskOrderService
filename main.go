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
	"userService/internal/transport/grpc"
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

	logger := log.New(os.Stdout, "Error: ", log.Lshortfile | log.Ltime)

	db, err := postgres.NewPostgresConnection(config)
	if err != nil {
		logger.Println(err)
		return
	}
	defer db.Close()

	userDB := postgresservice.NewPostgresService(db)
	userController := usercontroller.NewUserController(userDB)
	grpcController, err := grpc.NewGrpcClient(os.Getenv("GRPC_URL"), os.Getenv("GRPC_PORT"))
	if err != nil {
		logger.Println(err)
		return
	}

	handler := rest.NewHandler(userController, tokenManager, grpcController, logger)

	RunServer(fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT")), handler.InitRouter())
}

func RunServer(addr string, h http.Handler) {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	log.Println("waiting for connections...")
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	srv.Shutdown(ctx)
	log.Println("Server exited")
}