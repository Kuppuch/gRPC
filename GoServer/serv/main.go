package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "server/Protos/google.golang.org/grpc/greet"
	"sync"
	"time"
)

type server struct {
	pb.UnimplementedGreeterServer
}

var number = make(chan rxgo.Item)
var observable rxgo.Observable
var mutex = sync.Mutex{}
var cond = sync.NewCond(&mutex)

func main() {

	observable = rxgo.FromEventSource(number, rxgo.WithBackPressureStrategy(rxgo.Drop))

	go generate()

	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer(grpc.WriteBufferSize(0),grpc.ReadBufferSize(0), grpc.InitialConnWindowSize(1024))
	pb.RegisterGreeterServer(srv, &server{})
	fmt.Println("Serving")
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var counter int32 = 0

func (s *server) GetRandNum(req *pb.NumRequest, resp pb.Greeter_GetRandNumServer) error {
	fmt.Println("Метод вызван")

	for {
		mutex.Lock()
		cond.Wait()
		v := counter
		mutex.Unlock()
		err := resp.Send(&pb.NumReply{Message: v})
		fmt.Printf(" %v\n", v)
		//time.Sleep(5000 * time.Millisecond)
		if err != nil {
			fmt.Println("Client already disconnect")

			return err
		}
	}

}

func (s *server) GetSoloNum(context.Context, *pb.NumRequest) (*pb.NumReply, error) {
	mutex.Lock()
	cond.Wait()
	v := counter
	mutex.Unlock()
	return &pb.NumReply{Message: v}, nil

	return &pb.NumReply{}, nil
}

func generate() {
	for {
		mutex.Lock()
		counter++
		fmt.Printf("Generate %v\n", counter)
		cond.Broadcast()
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}
