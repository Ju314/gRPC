package main

import (
	"context"
	"gRPC/gRPC/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("Message received ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}
func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := "0.0.0.0:50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpc_Err := grpcServer.Serve(lis)

	if grpc_Err != nil {
		log.Fatal(grpc_Err)
	}
}
