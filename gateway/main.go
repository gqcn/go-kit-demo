package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// AddService 把两个东西加到一起
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
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

type addService struct{}

const maxLen = 10

var (
	// ErrTwoZeroes  Sum方法的业务规则不能对两个0求和
	ErrTwoZeroes = errors.New("can't sum two zeroes")

	// ErrIntOverflow Sum参数越界
	ErrIntOverflow = errors.New("integer overflow")

	// ErrTwoEmptyStrings Concat方法业务规则规定参数不能是两个空字符串.
	ErrTwoEmptyStrings = errors.New("can't concat two empty strings")

	// ErrMaxSizeExceeded Concat方法的参数超出范围
	ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)

// Sum 对两个数字求和，实现AddService。
func (s addService) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (math.MaxInt-b)) || (b < 0 && a < (math.MinInt-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

// Concat 连接两个字符串，实现AddService。
func (s addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrTwoEmptyStrings
	}
	if len(a)+len(b) > maxLen {
		return "", ErrMaxSizeExceeded
	}
	return a + b, nil
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ConcatRequest
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

func main() {
	svc := addService{}

	sumHandler := httptransport.NewServer(
		makeSumEndpoint(svc),
		decodeSumRequest,
		encodeResponse,
	)

	concatHandler := httptransport.NewServer(
		makeConcatEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/sum", sumHandler)
	http.Handle("/concat", concatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
