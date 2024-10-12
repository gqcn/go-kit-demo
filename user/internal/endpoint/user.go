package endpoint

import (
	"context"

	"go-kit-demo/user/api/user/v1"
	"go-kit-demo/user/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserEndpoint struct {
	v1.UnimplementedUserServer
	userService service.UserService
}

func NewUserEndpoint(ctx context.Context, client *mongo.Client) *UserEndpoint {
	return &UserEndpoint{
		userService: service.NewUserService(ctx, client),
	}
}

// Create 新增用户信息。
func (s *UserEndpoint) Create(ctx context.Context, req *v1.CreateRequest) (res *v1.EmptyResponse, err error) {
	return s.userService.Create(ctx, req)
}

// Search 查询符合条件的用户列表。
func (s *UserEndpoint) Search(ctx context.Context, req *v1.SearchRequest) (res *v1.SearchResponse, err error) {
	return s.userService.Search(ctx, req)
}
