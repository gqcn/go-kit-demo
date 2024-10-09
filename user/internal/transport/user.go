package transport

import (
	"context"

	"go-kit-demo/user/api"
	"go-kit-demo/user/internal/service"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	api.UnimplementedUserServer
	create grpctransport.Handler
	search grpctransport.Handler
}

func NewAddServer(svc *service.UserService) api.UserServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			makeCreateEndpoint(svc),
			directGRPCDecoder,
			directGRPCEncoder,
		),
		search: grpctransport.NewServer(
			makeSearchEndpoint(svc),
			directGRPCDecoder,
			directGRPCEncoder,
		),
	}
}

func directGRPCDecoder(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func directGRPCEncoder(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}

func makeCreateEndpoint(svc *service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return svc.Create(ctx, req.(*api.CreateRequest))
	}
}

func makeSearchEndpoint(svc *service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return svc.Search(ctx, req.(*api.SearchRequest))
	}
}
