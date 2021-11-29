Protocol Buffers - Google's data interchange format
Copyright 2008 Google Inc.
https://developers.google.com/protocol-buffers/

This package contains a precompiled binary version of the protocol buffer
compiler (protoc). This binary is intended for users who want to use Protocol
Buffers in languages other than C++ but do not want to compile protoc
themselves. To install, simply place this binary somewhere in your PATH.

If you intend to use the included well known types then don't forget to
copy the contents of the 'include' directory somewhere as well, for example
into '/usr/local/include/'.

Please refer to our official github site for more installation instructions:
  https://github.com/protocolbuffers/protobuf


需要protoc proto-gen-go proto-gen-go-grpc bin工具

protoc --go_out=. --govalidators_out=. hello.proto
protoc --go-grpc_out=. --govalidators_out=. hello.proto


protoc --proto_path=. --proto_path=/home/chenchong/gopath/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2/validator.proto --govalidators_out=. --go_out=. helloworld.proto


protoc  \
    --proto_path=./include \
    --proto_path=${GOPATH}/pkg/mod \
    --proto_path=${GOPATH}/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2 \
    --proto_path=. \
    --govalidators_out=. \
    --go-grpc_out=. \
    --go_out=. \
    *.proto



grpcurl grpc客户端

查看服务列表
grpcurl -plaintext 127.0.0.1:50051 list

查看某个服务的方法列表
grpcurl -plaintext 127.0.0.1:50051 list proto.Greeter

查看方法定义
grpcurl -plaintext 127.0.0.1:50051 describe proto.Greeter.SayHello

查看请求参数
grpcurl -plaintext 127.0.0.1:50051 describe proto.HelloRequest

请求服务
grpcurl -d '{"name": "zhangsan"}' -plaintext 127.0.0.1:50051 proto.Greeter.SayHello

