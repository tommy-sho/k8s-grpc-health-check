package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	proto "github.com/tommy-sho/grpc-loadbalncing/client/genproto"
	"google.golang.org/grpc"
)

var (
	gateway string
	port    string
	method  string
)

func init() {
	flag.StringVar(&gateway, "gateway", "gateway:50000", "gateway port")
	flag.StringVar(&method, "method", "Greeting", "method name")
}

func main() {

	flag.Parse()

	ctx := context.Background()

	gConn, err := grpc.DialContext(ctx, gateway, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed to connect with backend server error : %v ", err))
	}
	c := proto.NewGreetingServerClient(gConn)
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

L:
	for s.Scan() {
		n := s.Text()
		switch n {
		case "exit":
			break L
		default:
			r, err := c.Greeting(ctx, &proto.GreetingRequest{Name: n})
			if err != nil {
				fmt.Println("failed to call Gateway error : ", err)
				break L
			}
			fmt.Println(r.Message)
		}
		fmt.Print("> ")
	}
}
