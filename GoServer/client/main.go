package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "server/Protos/google.golang.org/grpc/greet"
	"time"
)

const (
	address = "0.0.0.0:5001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithReadBufferSize(0))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	stream, err := client.GetRandNum(context.Background() , &pb.NumRequest{}, grpc.MaxRetryRPCBufferSize(0))

	if err != nil {
		fmt.Println("Ошибка получения данных")
	}
	for {
		feature, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("Сервер завершил трансляцию данных")
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		fmt.Println(feature)
		time.Sleep(time.Second * 2)
	}
}