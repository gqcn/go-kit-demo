package service

import (
	"context"

	"go-kit-demo/user/api"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s UserService) Create(ctx context.Context, req *api.CreateRequest) (res *api.EmptyResponse, err error) {

	return nil, nil
}

func (s UserService) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {

	return nil, nil
}
