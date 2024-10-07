package routes

import (
	"src/index/src/controllers"
	pb "src/index/src/gRPC/hotel"

	"google.golang.org/grpc"
)

func InitHotelRoutes(s *grpc.Server) {
	pb.RegisterHotelServiceServer(s, &controllers.HotelServer{})
}
