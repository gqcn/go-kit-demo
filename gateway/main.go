package main

import (
	"context"
	"net/http"

	"go-kit-demo/gateway/internal/endpoint"
	"go-kit-demo/gateway/internal/transport"

	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		ctx         = context.Background()
		logger      = g.Log()
		config      = g.Cfg()
		userSvcAddr = config.MustGetWithEnv(ctx, `service.user`).String()
		serverAddr  = config.MustGetWithEnv(ctx, "server.address").String()
	)
	logger.Debugf(ctx, `user service addr: %s`, userSvcAddr)

	userClientConn, err := grpc.NewClient(
		userSvcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Fatalf(ctx, "did not connect: %v", err)
	}
	defer userClientConn.Close()

	var (
		svc           = endpoint.NewUserEndpoint(userClientConn)
		createHandler = transport.NewCreateHandler(svc)
		searchHandler = transport.NewSearchHandler(svc)
	)
	http.Handle("/create", createHandler)
	http.Handle("/search", searchHandler)
	logger.Infof(ctx, `gateway server starts listening on: "%s"`, serverAddr)
	logger.Fatal(ctx, http.ListenAndServe(serverAddr, nil))
}
