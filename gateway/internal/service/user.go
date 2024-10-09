package service

import (
	"context"

	"go-kit-demo/gateway/api"
	userSvc "go-kit-demo/user/api"
)

type UserService struct {
	userSvc *userSvc.UserClient
}

func NewAddService(userSvcClient *userSvc.UserClient) *UserService {
	return &UserService{
		userSvc: userSvcClient,
	}
}

func (s UserService) Create(ctx context.Context, req *api.CreateRequest) (res *api.CreateResponse, err error) {
	return nil, nil
}

func (s UserService) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {

	return nil, nil
}
