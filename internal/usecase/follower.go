package usecase

import (
	"context"
	"project/internal/models"
)

type FollowUseCase interface {
	GetAll(ctx context.Context, userId string) (*[]models.Follower, error)
	Follow(ctx context.Context, userId, followingId string) error
}

type FollowerInfoUseCase interface {
	GetFollower(ctx context.Context, userId string) (*[]models.FollowerInfo, error)
}

type FollowingInfoUseCase interface {
	GetFollowing(ctx context.Context, userId string) (*[]models.FollowingInfo, error)
}
