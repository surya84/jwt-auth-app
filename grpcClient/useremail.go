package grpcclient

import (
	"context"
	custommodels "jwtgen/customModels"
	pb "jwtgen/proto"
	"log"
	"time"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60000)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}
	log.Printf("%s", res.Message)
}

func GetData(client pb.GreetServiceClient, email string) (*custommodels.Employee, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60000)
	defer cancel()
	data := &pb.EmpRequest{
		Email: email,
	}
	res, err := client.GetEmpDetails(ctx, data)
	if err != nil {
		log.Println("could not response : %v", err)
		return nil, err
	}
	return &custommodels.Employee{
		Name:  res.Response.Email,
		EmpID: res.Response.EmpId,
		Email: res.Response.Email,
	}, nil
}
