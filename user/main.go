package main

import (
	"context"
	"net"

	"go-kit-demo/user/api"
	"go-kit-demo/user/internal/endpoint"
	"go-kit-demo/user/internal/transport"

	"github.com/gogf/gf/v2/frame/g"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	var (
		ctx        = context.Background()
		logger     = g.Log()
		config     = g.Cfg()
		mongoUri   = config.MustGetWithEnv(ctx, "mongodb.uri").String()
		serverAddr = config.MustGetWithEnv(ctx, "server.address").String()
		server     = grpc.NewServer()
	)
	// 初始化mongodb数据库客户端
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoUri),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// 初始化业务模块
	var (
		userService = endpoint.NewUserEndpoint(ctx, client)
		userServer  = transport.NewUserServer(userService)
	)
	logger.Infof(ctx, `grpc starts listening on: "%s"`, serverAddr)
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		logger.Fatalf(ctx, "failed to listen: %v", err)
	}
	api.RegisterUserServer(server, userServer)

	if err = server.Serve(listener); err != nil {
		logger.Fatalf(ctx, "failed to serve: %v", err)
	}
}
