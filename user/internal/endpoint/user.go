package endpoint

import (
	"context"

	"go-kit-demo/user/api"
	"go-kit-demo/user/internal/model"
	"go-kit-demo/user/internal/repository"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEndpoint struct {
	api.UnimplementedUserServer
	userRepo repository.UserRepository
}

func NewUserEndpoint(ctx context.Context, client *mongo.Client) *UserEndpoint {
	return &UserEndpoint{
		userRepo: repository.NewLocalUserRepository(ctx, client),
	}
}

// Create 新增用户信息。
func (s *UserEndpoint) Create(ctx context.Context, req *api.CreateRequest) (res *api.EmptyResponse, err error) {
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
	res = &api.EmptyResponse{}
	return
}

// Search 查询符合条件的用户列表。
// TODO 演示项目，未做分页。
func (s *UserEndpoint) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {
	g.Log().Debugf(ctx, `Search req: %s`, gjson.MustEncodeString(req))
	var (
		filter = bson.D{}
		fields = s.userRepo.CollectionInfo(ctx).Fields
	)
	// 检索条件判断
	if req.Gender > 0 {
		filter = append(filter, bson.E{Key: fields.Gender, Value: int(req.Gender)})
	}
	if req.MinAge > 0 {
		filter = append(filter, bson.E{Key: fields.Age, Value: bson.M{"$gte": int(req.MinAge)}})
	}
	if req.MaxAge > 0 {
		filter = append(filter, bson.E{Key: fields.Age, Value: bson.M{"$lte": int(req.MaxAge)}})
	}
	if req.MatchGender > 0 {
		filter = append(filter, bson.E{Key: fields.MatchGender, Value: req.MatchGender})
	}
	if req.MatchMinAge > 0 {
		filter = append(filter, bson.E{Key: fields.MatchMinAge, Value: bson.M{"$gte": int(req.MatchMinAge)}})
	}
	if req.MatchMaxAge > 0 {
		filter = append(filter, bson.E{Key: fields.MatchMaxAge, Value: bson.M{"$lte": int(req.MatchMaxAge)}})
	}
	// 执行数据库检索
	collection := s.userRepo.Collection(ctx)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, `search users failed`)
	}
	defer cur.Close(ctx)
	res = &api.SearchResponse{
		Users: make([]*api.UserData, 0),
	}
	// 查询数据到实体对象中
	var result = make([]model.User, 0)
	if err = cur.All(ctx, &result); err != nil {
		return nil, errors.Wrap(err, `mongodb scan result failed`)
	}
	// 将实体对象数据结构赋值到grpc返回数据结构中
	if err = gconv.Scan(result, &res.Users); err != nil {
		return nil, errors.Wrap(err, `result to users failed`)
	}
	return
}
