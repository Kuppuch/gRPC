package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "server/Protos/google.golang.org/grpc/greet"
	"sync"
	"time"
)

const (
	address = "0.0.0.0:5001"
)
var delay = flag.Duration("delay",100*time.Millisecond,"Delay between each receive")
var mutex = sync.Mutex{}
var cond = sync.NewCond(&mutex)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithReadBufferSize(0),
		grpc.WithWriteBufferSize(0))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	/*stream, err := client.GetRandNum(context.Background() , &pb.NumRequest{})
	if err != nil {
		fmt.Println("Ошибка получения данных")
	}*/

	for{
		repl, err := client.GetSoloNum(context.Background(), &pb.NumRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(repl.GetMessage())
		time.Sleep(*delay)
	}

	/*var va int32
	go func() {
		for {
			mutex.Lock()
			feature,err := stream.Recv()
			if err != nil {
				fmt.Println(err.Error(), "Нет данных")
			}
			if err == io.EOF {
				fmt.Println("Сервер завершил трансляцию данных")
				return
			}
			if err != nil {
				log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
			}
			va = feature.Message
			cond.Broadcast()
			mutex.Unlock()
		}
	}()

	for {

		mutex.Lock()
		cond.Wait()
		fmt.Println(va)
		mutex.Unlock()
		time.Sleep(*delay)
	}*/
}