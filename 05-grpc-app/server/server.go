package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
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
