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
	"userService/internal/transport/rest"
	usercontroller "userService/internal/user_controller"

	"github.com/joho/godotenv"
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

	db, err := postgres.NewPostgresConnection(config)
	if err != nil {
		log.Println("Error: database connect error")
		return
	}

	userDB := postgresservice.NewPostgresService(db)
	userController := usercontroller.NewUserController(userDB)

	handler := rest.NewHandler(userController, tokenManager)

	RunServer(fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT")), handler.InitRouter())
}

func RunServer(addr string, h http.Handler) {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	srv.Shutdown(ctx)
	log.Println("Server exited")
}
