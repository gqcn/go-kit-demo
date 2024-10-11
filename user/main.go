package main

import (
	"context"
	"net"

	"go-kit-demo/user/internal/transport"

	"github.com/gogf/gf/v2/frame/g"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// 所有的依赖在初始化时完成，也可以封装到独立的包中进行
	var (
		ctx        = context.Background()
		logger     = g.Log()
		config     = g.Cfg()
		mongoUri   = config.MustGetWithEnv(ctx, "mongodb.uri").String()
		serverAddr = config.MustGetWithEnv(ctx, "server.address").String()
		server     = grpc.NewServer()
	)
	// 初始化mongodb数据库客户端
	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoUri),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// 初始化业务模块
	transport.RegisterServerForUser(ctx, server, mongoClient)

	// 启动GRPC服务
	logger.Infof(ctx, `grpc starts listening on: "%s"`, serverAddr)
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		logger.Fatalf(ctx, "failed to listen: %v", err)
	}

	if err = server.Serve(listener); err != nil {
		logger.Fatalf(ctx, "failed to serve: %v", err)
	}
}
