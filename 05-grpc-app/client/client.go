package main

import (
	"context"
	"fmt"
	"io"
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

	//doRequestResponse(ctx, client)
	doServerStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
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

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	primeRequest := &proto.PrimeRequest{
		Start: 5,
		End:   100,
	}

	clientStream, err := client.GeneratePrimes(ctx, primeRequest)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		response, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("Thats all folks!")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(response.GetPrime())
	}
}
