package main

import pb "src/index/src/gRPC/user"

type Server struct {
	pb.UnimplementedUserServiceServer
}