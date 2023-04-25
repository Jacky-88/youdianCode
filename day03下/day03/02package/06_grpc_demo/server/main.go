package main

import (
	"day03/02package/06_grpc_demo/pb"
	"day03/02package/06_grpc_demo/serviceImpl"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//new一个grpc服务
	srv := grpc.NewServer()
	//注册服务
	pb.RegisterMessageSenderServer(srv, serviceImpl.MessageSenderServerImpl{})
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic("tcp启动失败" + err.Error())
	}
	err = srv.Serve(listener)
	if err != nil {
		panic("grpc启动失败" + err.Error())
	}
}


