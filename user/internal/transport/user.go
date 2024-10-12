package transport

import (
	"context"

	"go-kit-demo/user/api/user/v1"
	userep "go-kit-demo/user/internal/endpoint"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func RegisterServerForUser(ctx context.Context, server *grpc.Server, mongoClient *mongo.Client) {
	var userEndpoint = userep.NewUserEndpoint(ctx, mongoClient)
	v1.RegisterUserServer(server, userEndpoint)
}
