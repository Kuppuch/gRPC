package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	pb "server/Protos/google.golang.org/grpc/greet"
	"time"
)

type server struct {
	pb.UnimplementedGreeterServer
}

var b bool
var number = make(chan int32)

func main()  {

	//go generate(number)
	go func() {
		for true {
			val:= rand.Int31n(100)
			fmt.Println(val)
			b = true
			number <- val
			time.Sleep(time.Second)
		}
		close(number)
	}()


	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{})
	fmt.Println("Serving")
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetRandNum(req *pb.NumRequest, resp pb.Greeter_GetRandNumServer) error {
	for true {
		if b {
			val:=<-number
			fmt.Printf("Send client %v",val)
			err := resp.Send(&pb.NumReply{Message: val})
			if err!=nil {
				return err
			}
			b = false
		}
	}
	return nil
}

func generate(number chan int){
	for {
		number <- rand.Intn(100)
		b = true
		time.Sleep(time.Second)
	}

	close(number)
}