package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/tommy-sho/k8s-grpc-health-check/proto"
	"google.golang.org/grpc"
)

var (
	gateway string
	name    string
)

func init() {
	flag.StringVar(&gateway, "gateway", "localhost:31000", "gateway address")
	flag.StringVar(&name, "name", "", "input your name!")
}

func main() {

	flag.Parse()

	ctx := context.Background()

	gConn, err := grpc.DialContext(ctx, gateway, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed to connect with backend server error : %v ", err))
	}
	c := proto.NewGatewayServerClient(gConn)

	r, err := c.Greeting(ctx, &proto.GreetingRequest{Name: name})
	if err != nil {
		fmt.Println("failed to call Gateway error : ", err)
		return
	}
	fmt.Println(r.Message)
}
