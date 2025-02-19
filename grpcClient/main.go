package grpcclient

import (
	pb "jwtgen/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func GrpcClinet() pb.GreetServiceClient {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while connecting to server .. %v", err)
	}
	client := pb.NewGreetServiceClient(conn)
	CallSayHello(client)
	return client
}
