package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc/pb"
)

type server struct {
	pb.UnimplementedGreeterServer // 必须内嵌这个结构体
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 1.服务器监听端口
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	// 2.初始化grpc服务
	srv := grpc.NewServer()
	// 3.将实现的服务器接口的结构体注册到grpc服务中
	pb.RegisterGreeterServer(srv, &server{})
	// 4.启动grpc服务
	if err := srv.Serve(lis); err != nil {
		log.Fatal("fail")
	}
}