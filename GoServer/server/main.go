package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
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
var number = make(chan rxgo.Item)
var observable rxgo.Observable

func main()  {
	observable = rxgo.FromEventSource(number, rxgo.WithBackPressureStrategy(rxgo.Drop))

	go generate(number)

	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer(grpc.ReadBufferSize(0),grpc.WriteBufferSize(0))
	pb.RegisterGreeterServer(srv, &server{})
	fmt.Println("Serving")
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetRandNum(req *pb.NumRequest, resp pb.Greeter_GetRandNumServer) error {
	fmt.Println("Метод вызван")
	dataChannel:=observable.Observe()
	//nc:=make( <- chan int32)
	//close(nc)
	//close(dataChannel)
	err := resp.Context().Err()
	if err!=nil{
		return err
	}
	for {
		select {
			case <-resp.Context().Done():
				fmt.Println("Client disconnect")
				return resp.Context().Err()
			case va:=<- dataChannel:
				fmt.Println(" Send ", va)
				err := resp.Send(&pb.NumReply{Message: va.V.(int32)})
				if err!=nil {
					fmt.Println("Client already disconnect")
					return err
				}
		}
	}
}

func generate(number chan rxgo.Item) {
	//var val int32 = 0
	for {
		val := rand.Int31n(100)
		//val++
		number <- rxgo.Item{V: val}
		time.Sleep(time.Second)
	}
}