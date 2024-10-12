package endpoint

import (
	"context"

	"go-kit-demo/gateway/api/user/v1"
	"go-kit-demo/gateway/internal/service"

	"google.golang.org/grpc"
)

type UserEndpoint struct {
	userSvc service.UserService
}

func NewUserEndpoint(userClientConn *grpc.ClientConn) *UserEndpoint {
	return &UserEndpoint{
		userSvc: service.NewUserService(userClientConn),
	}
}

// Create 创建用户请求。
func (s *UserEndpoint) Create(ctx context.Context, req *v1.CreateRequest) (res *v1.CreateResponse, err error) {
	return s.userSvc.Create(ctx, req)
}

// Search 查询用户列表接口。
func (s *UserEndpoint) Search(ctx context.Context, req *v1.SearchRequest) (res *v1.SearchResponse, err error) {
	return s.userSvc.Search(ctx, req)
}
