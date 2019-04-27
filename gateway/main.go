package main

import (
	"context"
	"fmt"
	"net"
	"os"

	proto "github.com/tommy-sho/k8s-grpc-health-check/gateway/genproto"
	"github.com/tommy-sho/k8s-grpc-health-check/gateway/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50002"
)

func main() {
	ctx := context.Background()

	bConn, err := grpc.DialContext(ctx, os.Getenv("BACKEND_PORT"), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed to connect with backend server error : %v ", err))
	}
	bClient := proto.NewBackendServerClient(bConn)
	g := server.NewGatewaySerive(bClient)
	s := grpc.NewServer()

	proto.RegisterGreetingServerServer(s, g)
	server.RegisterHeathCheck(s)
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
