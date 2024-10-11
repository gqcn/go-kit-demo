package endpoint

import (
	"context"

	"go-kit-demo/gateway/api"
	userPb "go-kit-demo/user/api"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type UserEndpoint struct {
	userSvc userPb.UserClient
}

func NewUserEndpoint(userClientConn *grpc.ClientConn) *UserEndpoint {
	return &UserEndpoint{
		userSvc: userPb.NewUserClient(userClientConn),
	}
}

// Create 创建用户请求。
func (s *UserEndpoint) Create(ctx context.Context, req *api.CreateRequest) (res *api.CreateResponse, err error) {
	// 数据校验
	if err = g.Validator().Data(req).Run(ctx); err != nil {
		return
	}
	g.Log().Debugf(ctx, `Create req: %s`, gjson.MustEncodeString(req))
	// http层数据结构转换为grpc请求数据格式
	var createReq = userPb.CreateRequest{}
	if err = gconv.Scan(req, &createReq); err != nil {
		return
	}
	if _, err = s.userSvc.Create(ctx, &createReq); err != nil {
		return nil, errors.Wrap(err, "create user failed")
	}
	res = &api.CreateResponse{}
	return
}

// Search 查询用户列表接口。
// TODO 演示项目，未做分页。
func (s *UserEndpoint) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {
	// 数据校验
	if err = g.Validator().Data(req).Run(ctx); err != nil {
		return
	}
	g.Log().Debugf(ctx, `Search req: %s`, gjson.MustEncodeString(req))
	var (
		searchReq userPb.SearchRequest
		searchRes *userPb.SearchResponse
	)
	// http层数据结构转换为grpc请求数据格式
	if err = gconv.Scan(req, &searchReq); err != nil {
		return
	}
	if searchRes, err = s.userSvc.Search(ctx, &searchReq); err != nil {
		return nil, errors.Wrap(err, `search users failed`)
	}
	res = &api.SearchResponse{}
	if err = gconv.Scan(searchRes.Users, &res.Users); err != nil {
		return
	}
	return
}
