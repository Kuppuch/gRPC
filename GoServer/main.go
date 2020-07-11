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
	/*go func() {
		var val int32
		for true {

			val = rand.Int31n(100)
			number <- val
			//number <- rand.Int31n(100)
			fmt.Println(" Г", val)
			//i++
			//b = true
			//number <- val
			time.Sleep(time.Second)

		}
		close(number)
	}()*/

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

	fmt.Println("Метод вызван")
	for true {
		for item := range observable.Observe() {
			va := item.V
			fmt.Println(" Send ", va)
			err := resp.Send(&pb.NumReply{Message: va.(int32)})
			if err!=nil {
				return err
			}
		}
	}
	return nil
}

func generate(number chan rxgo.Item) {
	for {
		val := rand.Int31n(100)
		number <- rxgo.Item{V: val}
		time.Sleep(time.Second * 2)
	}
}