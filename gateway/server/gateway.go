package server

import (
	"context"
	"fmt"
	"os"

	proto "github.com/tommy-sho/grpc-loadbalncing/gateway/genproto"
)

type GatewayService struct {
	backendClient proto.BackendServerClient
}

func NewGatewaySerive(bClient proto.BackendServerClient) *GatewayService {
	return &GatewayService{bClient}
}

func (g *GatewayService) Greeting(ctx context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	in := &proto.MessageRequest{
		Name: req.Name,
	}
	r, err := g.backendClient.Message(ctx, in)
	if err != nil {
		return &proto.GreetingResponse{}, fmt.Errorf("gateway backendClinet error : %v ", err)
	}

	res := &proto.GreetingResponse{
		Message: fmt.Sprintf("%v, gateway IP: %v ", r.Message, os.Getenv("MY_POD_IP")),
	}
	return res, nil
}
