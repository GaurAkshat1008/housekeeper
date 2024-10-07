package common

import (
	pb "src/index/src/gRPC/user"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}