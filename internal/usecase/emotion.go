package usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeUseCase interface {
	GetLike(ctx context.Context, userId primitive.ObjectID) ([]any, error)
	Like(ctx context.Context, videoId, userId string) error
}

type LoveUseCase interface {
	GetLove(ctx context.Context, userId primitive.ObjectID) ([]any, error)
	Love(ctx context.Context, videoId, userId string) error
}
