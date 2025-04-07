package repositories

import (
	"project/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type FollowerRepo struct {
	GenericRepository[models.Follower]
}

func NewFollowerRepo(db *mongo.Database, collectionName string) *FollowerRepo {
	baseRepo := NewBaseRepo[models.Follower](db, collectionName)
	return &FollowerRepo{baseRepo}
}

type FollowerInfoRepo struct {
	GenericRepository[models.FollowerInfo]
}

func NewFollowerInfoRepo(db *mongo.Database, collectionName string) *FollowerInfoRepo {
	baseRepo := NewBaseRepo[models.FollowerInfo](db, collectionName)
	return &FollowerInfoRepo{baseRepo}
}

type FollowingInfoRepo struct {
	GenericRepository[models.FollowingInfo]
}

func NewFollowingInfoRepo(db *mongo.Database, collectionName string) *FollowingInfoRepo {
	baseRepo := NewBaseRepo[models.FollowingInfo](db, collectionName)
	return &FollowingInfoRepo{baseRepo}
}
