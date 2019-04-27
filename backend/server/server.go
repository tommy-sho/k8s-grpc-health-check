package server

import (
	"context"
	"fmt"

	proto "github.com/tommy-sho/grpc-loadbalncing/backend/genproto"
)

type BackendServer struct {
	userRepo BackendRepository
}

func NewBackendServer(userRepo BackendRepository) *BackendServer {
	return &BackendServer{
		userRepo: userRepo,
	}
}

func (b *BackendServer) Message(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	m, err := b.userRepo.GetMessageByName(req.Name)
	fmt.Println(req.Name)
	if err != nil {
		return &proto.MessageResponse{}, fmt.Errorf("Greeting error : %v ", err)
	}
	res := &proto.MessageResponse{
		Message: m,
	}
	return res, nil
}
