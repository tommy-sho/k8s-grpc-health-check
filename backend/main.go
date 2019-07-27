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

	"github.com/tommy-sho/k8s-grpc-health-check/lib"
	"github.com/tommy-sho/k8s-grpc-health-check/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50001"
)

func main() {
	b := &BackendServer{}
	server := grpc.NewServer()
	proto.RegisterBackendServerServer(server, b)
	lib.RegisterHeathCheck(server)
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

type BackendServer struct{}

func (b *BackendServer) Message(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	message := fmt.Sprintf("Hey! %s", req.Name)
	currentTime := time.Now()

	res := &proto.MessageResponse{
		Message:  message,
		Datetime: currentTime.Unix(),
	}
	return res, nil
}
