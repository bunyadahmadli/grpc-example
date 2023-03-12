package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	pb "grpc-example/gen/proto"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello World"})

	user, err := client.GetUser(context.Background(), &pb.UserRequest{Uuid: uuid.New().String()})
	if err != nil {
		return
	}
	fmt.Println(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
