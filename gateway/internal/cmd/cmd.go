package cmd

import (
	"context"
	"net/http"

	"go-kit-demo/gateway/internal/transport"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server for gateway",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				logger      = g.Log()
				config      = g.Cfg()
				router      = mux.NewRouter()
				userSvcAddr = config.MustGetWithEnv(ctx, `service.user`).String()
				serverAddr  = config.MustGetWithEnv(ctx, "server.address").String()
			)
			logger.Debugf(ctx, `configured user service addr: %s`, userSvcAddr)

			// user grpc connection.
			userClientConn, err := grpc.NewClient(
				userSvcAddr,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				logger.Fatalf(ctx, "create user grpc conn failed: %v", err)
			}
			defer userClientConn.Close()

			transport.RegisterRoutesForUser(ctx, router, userClientConn)

			logger.Infof(ctx, `gateway server starts listening on: "%s"`, serverAddr)
			if err := http.ListenAndServe(serverAddr, router); err != nil {
				logger.Warningf(ctx, `http server exit with error: %+v`, err)
			}
			return nil
		},
	}
)
