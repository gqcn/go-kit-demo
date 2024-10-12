package service

import (
	"context"

	v1 "go-kit-demo/user/api/user/v1"
	"go-kit-demo/user/internal/model"
	"go-kit-demo/user/internal/repository"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	v1.UserServer
}

type userServiceImpl struct {
	v1.UnimplementedUserServer
	userRepo repository.UserRepository
}

func NewUserService(ctx context.Context, client *mongo.Client) UserService {
	return &userServiceImpl{
		userRepo: repository.NewLocalUserRepository(ctx, client),
	}
}

// Create 新增用户信息。
func (s *userServiceImpl) Create(ctx context.Context, req *v1.CreateRequest) (res *v1.EmptyResponse, err error) {
	g.Log().Debugf(ctx, `Create req: %s`, gjson.MustEncodeString(req))
	// 将请求对象数据赋值到grpc的请求对象上
	var userModel model.User
	if err = gconv.Scan(req.User, &userModel); err != nil {
		return
	}
	// 调用GRPC创建用户
	if err = s.userRepo.Create(ctx, &userModel); err != nil {
		return
	}
	// 返回空对象，只要没有error就表示成功。
	// 根据业务场景定义返回数据结构。
	res = &v1.EmptyResponse{}
	return
}

// Search 查询符合条件的用户列表。
// TODO 演示项目，未做分页。
func (s *userServiceImpl) Search(ctx context.Context, req *v1.SearchRequest) (res *v1.SearchResponse, err error) {
	g.Log().Debugf(ctx, `Search req: %s`, gjson.MustEncodeString(req))
	var searchInput repository.SearchInput
	if err = gconv.Scan(req, &searchInput); err != nil {
		return nil, errors.Wrap(err, `scan to searchInput failed`)
	}
	result, err := s.userRepo.Search(ctx, searchInput)
	if err != nil {
		return nil, err
	}
	res = &v1.SearchResponse{
		Users: make([]*v1.UserData, 0),
	}
	// 将实体对象数据结构赋值到grpc返回数据结构中
	if err = gconv.Scan(result, &res.Users); err != nil {
		return nil, errors.Wrap(err, `result to users failed`)
	}
	return
}
