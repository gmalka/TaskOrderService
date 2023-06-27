package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"userService/auth/passwordManager"
	"userService/auth/tokenManager"
	"userService/pkg/database/postgres"
	postgresservice "userService/pkg/database/postgres/postgres_service"
	usercontroller "userService/pkg/user_controller"
	mygrpc "userService/transport/grpc"
	"userService/transport/rest"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	godotenv.Load()

	tokenManager := tokenManager.NewAuthService(os.Getenv("ACCESS_SECRET"), os.Getenv("REFRESH_SECRET"))

	config := postgres.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Sslmode:  os.Getenv("DB_SSLMODE"),
	}

	log := newLogger()

	db, err := postgres.NewPostgresConnection(config)
	if err != nil {
		log.Err.Println(err)
		return
	}
	defer db.Close()

	userDB := postgresservice.NewPostgresService(db)
	userController := usercontroller.NewUserController(userDB)

	// GRPC SETTINGS
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	path := fmt.Sprintf("%s:%s", os.Getenv("GRPC_URL"), os.Getenv("GRPC_PORT"))
	conn, err := grpc.Dial(path, opts...)
	if err != nil {
		log.Err.Printf("can't connect by grpc to path %s: %v", path, err)
		return
	}

	grpcController, err := mygrpc.NewGrpcClient(conn)
	if err != nil {
		log.Err.Println(err)
		return
	}

	handler := rest.NewHandler(userController, tokenManager, grpcController, passwordManager.NewPasswordManager(), log)

	RunServer(fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT")), handler.InitRouter(false), log)
}

func newLogger() rest.Log {

	loggerErr := log.New(os.Stderr, "ERROR:\t ", log.Lshortfile|log.Ltime)
	loggerInfo := log.New(os.Stdout, "INFO:\t ", log.Lshortfile|log.Ltime)

	log := rest.Log{
		Err: loggerErr,
		Inf: loggerInfo,
	}

	return log
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
