package main

import (
	"fmt"
	"net"

	pb "github.com/tommy-sho/grpc-loadbalncing/backend/genproto"
	"github.com/tommy-sho/grpc-loadbalncing/backend/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50001"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	r := server.NewUserRepository()
	g := server.NewBackendServer(r)

	s := grpc.NewServer()
	pb.RegisterBackendServerServer(s, g)
	if err != nil {
		panic(fmt.Errorf("new grpc server err: %v", err))
	}
	reflection.Register(s)

	s.Serve(lis)
}
