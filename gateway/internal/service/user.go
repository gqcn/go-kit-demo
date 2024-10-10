package service

import (
	"context"

	"go-kit-demo/gateway/api"
	userSvc "go-kit-demo/user/api"
)

type UserService interface {
	Create(ctx context.Context, req *api.CreateRequest) (res *api.CreateResponse, err error)
	Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error)
}

type localUserServiceImpl struct {
	userSvc *userSvc.UserClient
}

func NewAddService(userSvcClient *userSvc.UserClient) UserService {
	return &localUserServiceImpl{
		userSvc: userSvcClient,
	}
}

func (s localUserServiceImpl) Create(ctx context.Context, req *api.CreateRequest) (res *api.CreateResponse, err error) {
	return nil, nil
}

func (s localUserServiceImpl) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {
	return nil, nil
}
