package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	pb "server/Protos/google.golang.org/grpc/greet"
)

const (
	port = ":5001"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) GetRandNum(ctx context.Context, stream *pb.NumRequest) (*pb.NumReply, error) {
	for true {
		number := rand.Intn(100)
		log.Printf("Received: %v")
		return &pb.NumReply{Message: string(number)}, nil
	}

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
