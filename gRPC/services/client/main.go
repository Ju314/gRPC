package main

import (
	"context"
	"log"

	"gRPC/gRPC/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello gRPC",
	}

	res, err := client.RequestMessage(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.GetStatus())
}
