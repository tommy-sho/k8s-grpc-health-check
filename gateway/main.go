package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tommy-sho/k8s-grpc-health-check/lib"
	"github.com/tommy-sho/k8s-grpc-health-check/proto"
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

	b := &GatewayServer{client: bClient}
	server := grpc.NewServer()
	lib.RegisterHeathCheck(server)
	proto.RegisterGatewayServerServer(server, b)
	reflection.Register(server)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		<-stopChan
		gracefulStopChan := make(chan bool, 1)
		go func() {
			server.GracefulStop()
			gracefulStopChan <- true
		}()
		t := time.NewTimer(10 * time.Second)
		select {
		case <-gracefulStopChan:
			log.Print("Success graceful stop")
		case <-t.C:
			server.Stop()
		}
	}()

	errors := make(chan error)
	go func() {
		errors <- server.Serve(lis)
	}()

	if err := <-errors; err != nil {
		log.Fatal("Failed to server gRPC server", err)
	}

}

type GatewayServer struct {
	client proto.BackendServerClient
}

func (s *GatewayServer) Greeting(ctx context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	request := &proto.MessageRequest{
		Name: req.Name,
	}
	r, err := s.client.Message(ctx, request)
	if err != nil {
		return &proto.GreetingResponse{}, xerrors.Errorf("failed to request backendServer[name: %s]: %s", req.Name, err)
	}

	res := &proto.GreetingResponse{
		Message:  r.Message,
		Datetime: r.Datetime,
	}
	return res, nil
}
