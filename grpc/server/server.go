package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync/atomic"

	"google.golang.org/grpc"

	"grpc/pb"
)

type server struct {
	pb.UnimplementedGreeterServer // 必须内嵌这个结构体
}

var helloCount int32

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	newValue := atomic.AddInt32(&helloCount, 1)
	log.Printf("count is %d", newValue)
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

var helloStreamCount int32

func (s *server) HelloStream(stream pb.Greeter_HelloStreamServer) error {
	var r *pb.HelloRequest
	var err error
	for {
		r, err = stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("Received: %v", r.GetName())

		err = stream.Send(&pb.HelloReply{Message: "Hello " + r.GetName()})
		if err != nil {
			return err
		}
		newValue := atomic.AddInt32(&helloStreamCount, 1)
		log.Printf("count is %d", newValue)
	}
}

func main() {
	// 1.服务器监听端口
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	// 2.初始化grpc服务
	srv := grpc.NewServer()
	// 注册 grpcurl 所需的 reflection 服务
	//reflection.Register(server)
	// 3.将实现的服务器接口的结构体注册到grpc服务中
	pb.RegisterGreeterServer(srv, &server{})
	// 4.启动grpc服务
	if err := srv.Serve(lis); err != nil {
		log.Fatal("fail")
	}
}
