package service

import (
	"context"

	"go-kit-demo/user/api"
	"go-kit-demo/user/internal/model"
	"go-kit-demo/user/internal/repository"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Create(ctx context.Context, req *api.CreateRequest) (res *api.EmptyResponse, err error)
	Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error)
}

type localUserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(client *mongo.Client) UserService {
	return &localUserServiceImpl{
		userRepo: repository.NewLocalUserRepository(client),
	}
}

// Create 新增用户信息。
func (s *localUserServiceImpl) Create(ctx context.Context, req *api.CreateRequest) (res *api.EmptyResponse, err error) {
	err = s.userRepo.Create(ctx, &model.User{
		Name:   req.User.Name,
		Age:    int(req.User.Age),
		Gender: int(req.User.Gender),
		Location: &model.GeoPoint{
			Type:        req.User.Location.Type,
			Coordinates: req.User.Location.Coordinates,
		},
		MatchGender: int(req.User.MatchGender),
		MatchMinAge: int(req.User.MatchMinAge),
		MatchMaxAge: int(req.User.MatchMaxAge),
	})
	res = &api.EmptyResponse{}
	return
}

// Search 查询符合条件的用户列表。
// @TODO 演示场景，未做分页。
func (s *localUserServiceImpl) Search(ctx context.Context, req *api.SearchRequest) (res *api.SearchResponse, err error) {
	var (
		filter = bson.D{}
		fields = s.userRepo.CollectionInfo(ctx).Fields
	)
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
	collection := s.userRepo.Collection(ctx)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, `search users failed`)
	}
	defer cur.Close(ctx)
	res = &api.SearchResponse{
		Users: make([]*api.UserData, 0),
	}
	if err = cur.All(ctx, &res.Users); err != nil {
		return nil, errors.Wrap(err, `scan users failed`)
	}
	return
}
