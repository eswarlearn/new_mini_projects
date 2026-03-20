package main

import pb "github.com/eswar/50-grpc_golang/proto"

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponces, error) {
	return &pb.HellResponse{
		Message: "Hello",
	}, nil
}
