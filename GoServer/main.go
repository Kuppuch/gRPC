package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	pb "server/Protos/google.golang.org/grpc/greet"
)

type server struct {

}

func main()  {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{})
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetRandNum(*pb.NumRequest, pb.Greeter_GetRandNumServer) error{
	for true {
		number := rand.Intn(100)
		log.Printf("Received: %v")
		return &pb.NumReply{Message: string(number)}
	}
}

func (s *server) SayHelloSayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	number := rand.Intn(100)
	log.Printf("Received: %v")
	return &pb.HelloReply{Message: string(number)}, nil
}