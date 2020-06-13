package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go_msc/grpc/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) StreamSayHello(in *pb.HelloRequest, out pb.Greeter_StreamSayHelloServer) error {
	log.Printf("Received: %v", in.GetName())
	for i := 0; i < 4; i++ {
		out.Send(&pb.HelloReply{Message: fmt.Sprintf("Hello %d %v", i, in.GetName())})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("rpc server start")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
