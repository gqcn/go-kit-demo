package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"go-kit-demo/gateway/api"
	"go-kit-demo/gateway/internal/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewCreateHandler(svc *service.UserService) *httptransport.Server {
	return httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)
}

func NewSearchHandler(svc *service.UserService) *httptransport.Server {
	return httptransport.NewServer(
		makeSearchEndpoint(svc),
		decodeSearchRequest,
		encodeResponse,
	)
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeCreateEndpoint(svc *service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*api.CreateRequest)
		return svc.Create(ctx, req)
	}
}

func makeSearchEndpoint(svc *service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*api.SearchRequest)
		return svc.Search(ctx, req)
	}
}
