package repository

import (
	"context"

	"go-kit-demo/user/internal/consts"
	"go-kit-demo/user/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Collection(ctx context.Context) *mongo.Collection
	CollectionInfo(ctx context.Context) userCollectionInfo
	Create(ctx context.Context, user *model.User) error
}

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

func NewLocalUserRepository(client *mongo.Client) UserRepository {
	return &userRepositoryImpl{
		client:     client,
		collection: client.Database(consts.DBName).Collection(collectionName),
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

// CollectionInfo 返回数据集合信息。
func (u *userRepositoryImpl) CollectionInfo(_ context.Context) userCollectionInfo {
	return u.collectionInfo
}

// Collection 返回数据库用户集合对象，用于上层不可抽象的自定义数据库操作逻辑。
func (u *userRepositoryImpl) Collection(_ context.Context) *mongo.Collection {
	return u.collection
}

// Create 抽象封装创建用户的数据库操作逻辑。
func (u *userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	_, err := u.collection.InsertOne(ctx, user)
	return err
}
