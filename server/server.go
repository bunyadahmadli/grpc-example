package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-example/gen/proto"
	"log"
	"net"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	res := &pb.UserResponse{
		Name:  "deneme 1",
		Age:   19,
		Email: "deneme@deneme.com",
	}

	return res, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	// ...
	listener, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		return
	}

}
