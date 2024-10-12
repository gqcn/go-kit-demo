package repository

import (
	"context"

	"go-kit-demo/user/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository 对数据库操作进行抽象封装。
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Search(ctx context.Context, in SearchInput) ([]model.User, error)
}

// 数据集合信息，避免程序中对使用字段进行硬编码。
// 该信息可以通过工具自动生成，以保持和数据集合的同步，避免人工维护易造成的程序结构与数据集合结构的差异。
type userCollectionInfo struct {
	Name   string
	Fields userCollectionInfoFields
}

type userCollectionInfoFields struct {
	Name        string
	Age         string
	Gender      string
	Location    string
	MatchGender string
	MatchMinAge string
	MatchMaxAge string
}

type userRepositoryImpl struct {
	client         *mongo.Client
	collection     *mongo.Collection
	collectionInfo userCollectionInfo
}

const (
	collectionName = "user"
)

func NewLocalUserRepository(ctx context.Context, client *mongo.Client) UserRepository {
	var dbName = g.Cfg().MustGetWithEnv(ctx, "db_name").String()
	return &userRepositoryImpl{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
		collectionInfo: userCollectionInfo{
			Name: collectionName,
			Fields: userCollectionInfoFields{
				Name:        "name",
				Age:         "age",
				Gender:      "gender",
				Location:    "location",
				MatchGender: "match_gender",
				MatchMinAge: "match_min_age",
				MatchMaxAge: "match_max_age",
			},
		},
	}
}

// Create 抽象封装创建用户的数据库操作逻辑。
func (u *userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	_, err := u.collection.InsertOne(ctx, user)
	return err
}

type SearchInput struct {
	MinAge      int // 年龄范围(min)
	MaxAge      int // 年龄范围(max)
	Gender      int // 性别
	MatchGender int // 性取向
	MatchMinAge int // 年龄偏好(min)
	MatchMaxAge int // 年龄偏好(max)
}

func (u *userRepositoryImpl) Search(ctx context.Context, in SearchInput) ([]model.User, error) {
	var (
		filter = bson.D{}
		fields = u.collectionInfo.Fields
	)
	// 检索条件判断
	if in.Gender > 0 {
		filter = append(filter, bson.E{Key: fields.Gender, Value: in.Gender})
	}
	if in.MinAge > 0 {
		filter = append(filter, bson.E{Key: fields.Age, Value: bson.M{"$gte": in.MinAge}})
	}
	if in.MaxAge > 0 {
		filter = append(filter, bson.E{Key: fields.Age, Value: bson.M{"$lte": in.MaxAge}})
	}
	if in.MatchGender > 0 {
		filter = append(filter, bson.E{Key: fields.MatchGender, Value: in.MatchGender})
	}
	if in.MatchMinAge > 0 {
		filter = append(filter, bson.E{Key: fields.MatchMinAge, Value: bson.M{"$gte": in.MatchMinAge}})
	}
	if in.MatchMaxAge > 0 {
		filter = append(filter, bson.E{Key: fields.MatchMaxAge, Value: bson.M{"$lte": in.MatchMaxAge}})
	}
	// 执行数据库检索
	cur, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, `search users failed`)
	}
	defer cur.Close(ctx)
	// 查询数据到实体对象中
	var result = make([]model.User, 0)
	if err = cur.All(ctx, &result); err != nil {
		return nil, errors.Wrap(err, `mongodb scan result failed`)
	}
	return result, nil
}
