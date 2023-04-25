package serviceImpl

import (
	"context"
	"day03/02package/06_grpc_demo/pb"
	"fmt"
)

type MessageSenderServerImpl struct {
	*pb.UnimplementedMessageSenderServer
}

func (MessageSenderServerImpl) Send(context context.Context,
	request *pb.MessageRequest) (*pb.MessageResponse, error) {
		fmt.Println("server receive message:", request.GetSaySomething())
		resp := &pb.MessageResponse{}
		resp.ResponseSomething = "roger that!"
		return resp, nil
}