package main

import (
	"flag"
	"fmt"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"google.golang.org/grpc"
	"mxshop_srvs/user_srv/handler"
	"net"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "IP address")
	Port := flag.Int("port", 50051, "Port number")
	flag.Parse()

	fmt.Println("IP:", *IP, "Port:", *Port)
	server := grpc.NewServer()
	userProto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
