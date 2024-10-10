package transport

import (
	"context"

	"go-kit-demo/user/api"
	userep "go-kit-demo/user/internal/endpoint"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type userServer struct {
	api.UnimplementedUserServer
	create grpctransport.Handler
	search grpctransport.Handler
}

func NewUserServer(ep userep.UserEndpoint) api.UserServer {
	return &userServer{
		create: grpctransport.NewServer(
			makeCreateEndpoint(ep),
			directGRPCDecoder,
			directGRPCEncoder,
		),
		search: grpctransport.NewServer(
			makeSearchEndpoint(ep),
			directGRPCDecoder,
			directGRPCEncoder,
		),
	}
}

// Create 新增用户信息。
func (s *userServer) Create(ctx context.Context, req *api.CreateRequest) (res *api.EmptyResponse, err error) {
	_, response, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*api.EmptyResponse), nil
}

// Search 查询符合条件的用户列表。
// @TODO 演示场景，未做分页。
func (s *userServer) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {
	_, response, err := s.search.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*api.SearchResponse), nil
}

func directGRPCDecoder(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func directGRPCEncoder(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}

func makeCreateEndpoint(ep userep.UserEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return ep.Create(ctx, req.(*api.CreateRequest))
	}
}

func makeSearchEndpoint(ep userep.UserEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return ep.Search(ctx, req.(*api.SearchRequest))
	}
}
