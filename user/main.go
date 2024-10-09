package main

import (
	"fmt"
	"net"

	"go-kit-demo/user/api"
	"go-kit-demo/user/internal/service"
	"go-kit-demo/user/internal/transport"

	"google.golang.org/grpc"
)

func main() {
	var (
		svc = service.NewUserService()
		gs  = transport.NewAddServer(svc)
	)
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	api.RegisterUserServer(s, gs)
	if err = s.Serve(listener); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
