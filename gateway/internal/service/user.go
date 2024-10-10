package service

import (
	"context"

	"go-kit-demo/gateway/api"
	userPb "go-kit-demo/user/api"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type UserService interface {
	Create(ctx context.Context, req *api.CreateRequest) (res *api.CreateResponse, err error)
	Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error)
}

type userServiceImpl struct {
	userSvc userPb.UserClient
}

func NewUserService(clientConn *grpc.ClientConn) UserService {
	return &userServiceImpl{
		userSvc: userPb.NewUserClient(clientConn),
	}
}

func (s *userServiceImpl) Create(ctx context.Context, req *api.CreateRequest) (res *api.CreateResponse, err error) {
	var createReq = userPb.CreateRequest{}
	if err = gconv.Scan(req, &createReq.User); err != nil {
		return
	}
	_, err = s.userSvc.Create(ctx, &createReq)
	if err != nil {
		return nil, errors.Wrap(err, "create user failed")
	}
	res = &api.CreateResponse{}
	return
}

func (s *userServiceImpl) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {
	var searchReq = userPb.SearchRequest{}
	if err = gconv.Scan(req, &searchReq); err != nil {
		return
	}
	searchRes, err := s.userSvc.Search(ctx, &searchReq)
	if err != nil {
		return nil, errors.Wrap(err, `search users failed`)
	}
	res = &api.SearchResponse{}
	if err = gconv.Scan(searchRes.Users, &res.Users); err != nil {
		return
	}
	return
}
