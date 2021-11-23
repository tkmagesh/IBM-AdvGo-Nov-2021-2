package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

//implementation of the service
type server struct {
	proto.UnimplementedAppServiceServer
}

//implementation of the service operation
func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	resp := &proto.AddResponse{
		Sum: result,
	}
	return resp, nil
}

func (s *server) GeneratePrimes(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for i := start; i <= end; i++ {
		if isPrime(i) {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Sending Prime : %d\n", i)
			resp := &proto.PrimeResponse{
				Prime: i,
			}
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
	return nil
}

func isPrime(no int32) bool {
	var i int32
	for i = 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		no := req.GetNum()
		fmt.Println("Average Req received : ", no)
		sum += no
		count++
	}
	avg := sum / count
	resp := &proto.AverageResponse{
		Average: avg,
	}
	fmt.Println("Average response : ", avg)
	return serverStream.SendAndClose(resp)
}

func (s *server) CheckPrime(serverStream proto.AppService_CheckPrimeServer) error {
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		no := req.GetNum()
		isPrime := isPrime(no)
		time.Sleep(1 * time.Second)
		resp := &proto.IsPrimeResponse{
			Num:     no,
			IsPrime: isPrime,
		}
		if err := serverStream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	s := &server{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}
