package main

import (
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

func main()  {
	observable = rxgo.FromEventSource(number, rxgo.WithBackPressureStrategy(rxgo.Drop))

	//go generate(number)

	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer(grpc.WriteBufferSize(0))
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
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go generate(mutex, cond)

	for {
		mutex.Lock()
		cond.Wait()
		v := counter
		mutex.Unlock()
		err := resp.Send(&pb.NumReply{Message: v})
		fmt.Printf(" %v\n", v)
		//time.Sleep(5000 * time.Millisecond)
		if err!=nil {
			fmt.Println("Client already disconnect")

		}
	}

}

func generate(mutex sync.Mutex, cond *sync.Cond) {
	//var val int32 = 0
	for {
		/*val := rand.Int31n(100)
		//val++
		number <- rxgo.Item{V: val}
		time.Sleep(time.Second)*/
		mutex.Lock()
		counter++
		fmt.Printf("Generate %v\n", counter)
		cond.Broadcast()
		mutex.Unlock()
		time.Sleep(1000 * time.Millisecond)
	}
}
