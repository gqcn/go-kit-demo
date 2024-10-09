package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"go-kit-demo/gateway/api"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// AddService 把两个东西加到一起
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

func NewSumHandler(svc AddService) *httptransport.Server {
	return httptransport.NewServer(
		makeSumEndpoint(svc),
		decodeSumRequest,
		encodeResponse,
	)
}

func NewConcatHandler(svc AddService) *httptransport.Server {
	return httptransport.NewServer(
		makeConcatEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SumRequest)
		v, err := svc.Sum(ctx, req.A, req.B)
		if err != nil {
			return api.SumResponse{V: v, Err: err.Error()}, nil
		}
		return api.SumResponse{V: v}, nil
	}
}

func makeConcatEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ConcatRequest)
		v, err := svc.Concat(ctx, req.A, req.B)
		if err != nil {
			return api.ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return api.ConcatResponse{V: v}, nil
	}
}
