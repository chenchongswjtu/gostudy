package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpc/pb"
)

func main() {
	// 1.连接grpc服务地址
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2.根据连接创建一个grpc的客户端
	client := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "world"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 3.通过客户端远程调用函数
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	stream, err := client.HelloStream(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	err = stream.Send(&pb.HelloRequest{Name: "sss"})
	if err != nil {
		log.Fatalf("2222 %s", err)
	}

	recv, err := stream.Recv()
	if err != nil {
		log.Fatalf("3333 %s", err)
	}

	log.Println("client", recv.GetMessage())

	stream.CloseSend()
}
