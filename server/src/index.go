package main

import (
	"context"
	"fmt"
	"net"

	"src/index/src/db"
	"src/index/src/routes"
	"src/index/src/utils"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	s *grpc.Server
)

func initRoutes() {
	routes.InitUserRoutes(s)
	routes.InitHotelRoutes(s)
}

func init() {

	utils.InitLogger()
	err := godotenv.Load(".env.development")
	if err != nil {
		utils.Logger.Fatal("Error loading .env file")
	}

	if err != nil {
		fmt.Println(err)
		utils.Logger.Fatalf("Failed to connect to redis!")
	}
	s = grpc.NewServer(
		grpc.UnaryInterceptor(contextInterceptor()),
	)
	initRoutes()
	db.ConnectDatabase()
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		utils.Logger.Fatalf("Failed to listen: %v", err)
	}
	utils.Logger.Println("Server started at port :5000")
	if err := s.Serve(lis); err != nil {
		utils.Logger.Fatalf("Failed to serve: %v", err)
	}
}

func contextInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = context.Background()
		ctx, err := utils.InitRedisClient(ctx)
		if err != nil {
			utils.Logger.Fatalf("Failed to connect to redis!")
		}
		return handler(ctx, req)
	}
}
