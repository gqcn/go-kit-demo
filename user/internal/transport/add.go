package transport

import (
	"context"

	"go-kit-demo/user/api"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// AddService 把两个东西加到一起
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

type grpcServer struct {
	api.UnimplementedAddServer
	sum    grpctransport.Handler
	concat grpctransport.Handler
}

// SumRequest Sum方法的参数.
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// SumResponse Sum方法的响应
type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

// ConcatRequest Concat方法的参数.
type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

// ConcatResponse  Concat方法的响应.
type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func NewAddServer(svc AddService) api.AddServer {
	return &grpcServer{
		sum: grpctransport.NewServer(
			makeSumEndpoint(svc),
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
		concat: grpctransport.NewServer(
			makeConcatEndpoint(svc),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
		),
	}
}

// decodeGRPCSumRequest 将Sum方法的gRPC请求参数转为内部的SumRequest
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*api.SumRequest)
	return SumRequest{A: int(req.A), B: int(req.B)}, nil
}

// decodeGRPCConcatRequest 将Concat方法的gRPC请求参数转为内部的ConcatRequest
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*api.ConcatRequest)
	return ConcatRequest{A: req.A, B: req.B}, nil
}

// encodeGRPCSumResponse 封装Sum的gRPC响应
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(SumResponse)
	return &api.SumResponse{V: int64(resp.V), Err: resp.Err}, nil
}

// encodeGRPCConcatResponse 封装Concat的gRPC响应
func encodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ConcatResponse)
	return &api.ConcatResponse{V: resp.V, Err: resp.Err}, nil
}

func makeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := svc.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

func makeConcatEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		v, err := svc.Concat(ctx, req.A, req.B)
		if err != nil {
			return ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return ConcatResponse{V: v}, nil
	}
}

func (s *grpcServer) Sum(ctx context.Context, req *api.SumRequest) (*api.SumResponse, error) {
	_, rep, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*api.SumResponse), nil
}

func (s *grpcServer) Concat(ctx context.Context, req *api.ConcatRequest) (*api.ConcatResponse, error) {
	_, rep, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*api.ConcatResponse), nil
}
