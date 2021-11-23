package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"

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
