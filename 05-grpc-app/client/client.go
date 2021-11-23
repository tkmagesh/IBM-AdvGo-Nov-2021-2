package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//doServerStreaming(ctx, client)
	//doClientStreaming(ctx, client)
	doBiDiStreaming(ctx, client)
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

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	nos := []int32{5, 2, 6, 3, 1, 7, 9, 4, 8, 10}
	clientStream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Sending No : %d\n", no)
		req := &proto.AverageRequest{
			Num: no,
		}
		if er := clientStream.Send(req); er != nil {
			log.Fatalln(er)
		}
	}
	response, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average : ", response.GetAverage())
}

func doBiDiStreaming(ctx context.Context, client proto.AppServiceClient) {
	//sending the requests
	nos := []int32{5, 2, 6, 3, 1, 7, 9, 4, 8, 10}
	clientStream, err := client.CheckPrime(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan bool)
	go func() {
		for {
			response, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(response.GetNum(), response.GetIsPrime())
		}
		done <- true
	}()
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Sending No : %d\n", no)
		req := &proto.IsPrimeRequest{
			Num: no,
		}
		if er := clientStream.Send(req); er != nil {
			log.Fatalln(er)
		}
	}

	<-done
}
