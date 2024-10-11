package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"go-kit-demo/gateway/api"
	userep "go-kit-demo/gateway/internal/endpoint"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func RegisterRoutesForUser(_ context.Context, router *mux.Router, userClientConn *grpc.ClientConn) {
	// 用户相关接口路由注册
	var userEndpoint = userep.NewUserEndpoint(userClientConn)
	router.Methods("POST").Path("/create").Handler(httptransport.NewServer(
		makeCreateEndpoint(userEndpoint),
		decodeCreateRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/search").Handler(httptransport.NewServer(
		makeSearchEndpoint(userEndpoint),
		decodeSearchRequest,
		encodeResponse,
	))
}

func makeCreateEndpoint(ep *userep.UserEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*api.CreateRequest)
		return ep.Create(ctx, req)
	}
}

func makeSearchEndpoint(ep *userep.UserEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*api.SearchRequest)
		return ep.Search(ctx, req)
	}
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
