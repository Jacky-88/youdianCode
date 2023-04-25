package main

import (
	"context"
	"day03/02package/06_grpc_demo/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//建立grpc连接
	conn, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc连接失败" + err.Error())
	}
	defer conn.Close()
	//new一个 grpc client对象
	client := pb.NewMessageSenderClient(conn)
	resp, err := client.Send(context.Background(), &pb.MessageRequest{
		SaySomething: "hello grpc!",
	})
	if err != nil {
		panic("grpc发送失败" + err.Error())
	}
	fmt.Println("client receive message: ", resp.ResponseSomething)
}