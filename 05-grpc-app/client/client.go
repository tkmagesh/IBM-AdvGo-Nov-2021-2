package main

import (
	"context"
	"fmt"
	"log"

	"grpc-app/proto"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	//proxy instance
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	//sending the request
	response, err := client.Add(ctx, addRequest)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(response.GetSum())
}
