package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"go-kit-demo/gateway/api"
	userep "go-kit-demo/gateway/internal/endpoint"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewCreateHandler(ep userep.UserEndpoint) *httptransport.Server {
	return httptransport.NewServer(
		makeCreateEndpoint(ep),
		decodeCreateRequest,
		encodeResponse,
	)
}

func NewSearchHandler(ep userep.UserEndpoint) *httptransport.Server {
	return httptransport.NewServer(
		makeSearchEndpoint(ep),
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

func makeCreateEndpoint(ep userep.UserEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*api.CreateRequest)
		return ep.Create(ctx, req)
	}
}

func makeSearchEndpoint(ep userep.UserEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*api.SearchRequest)
		return ep.Search(ctx, req)
	}
}
