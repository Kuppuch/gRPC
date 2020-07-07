package main

import (
	//"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	pb "server/Protos/google.golang.org/grpc/greet"
	"time"
)

type server struct {

}

var number int
var b bool

func main()  {
	go generate()
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	//pb.RegisterGreeterServer(srv, &server{})
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetRandNum(req *pb.NumRequest, stream pb.Greeter_GetRandNumServer) *pb.NumReply {
	/*for true {
		number := rand.Intn(100)
		log.Printf("Received: %v")
		return &pb.NumReply{Message: string(number)}
	}
	return nil*/

	for true {
		if b {
			b = false
			return &pb.NumReply{Message: string(number)}
		}
	}
	return nil
}

/*func (s *server) SayHelloSayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	if b {
		b = false
		log.Printf("Received: %v")
		return &pb.HelloReply{Message: string(number)}, nil
	}

}*/

func generate(){
	for  true {
		number = rand.Intn(100)
		b = true
		fmt.Println(number)
		time.Sleep(time.Second)
	}
}