package routes

import (
	"src/index/src/controllers"
	pb "src/index/src/gRPC/user"

	"google.golang.org/grpc"
)

func InitUserRoutes(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, &controllers.UserServer{})
}
