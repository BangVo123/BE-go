package repositories

import (
	"project/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type VideoRepo struct {
	GenericRepository[models.Video]
}

func NewVideoRepo(db *mongo.Database, modelName string) *VideoRepo {
	BaseRepo := NewBaseRepo[models.Video](db, modelName)
	return &VideoRepo{GenericRepository: BaseRepo}
}

type VideoWithOwnerInfoRepo struct {
	GenericRepository[models.VideoWithOwnerInfo]
}

func NewVideoWithOwnerInfoRepo(db *mongo.Database, modelName string) *VideoWithOwnerInfoRepo {
	BaseRepo := NewBaseRepo[models.VideoWithOwnerInfo](db, modelName)
	return &VideoWithOwnerInfoRepo{GenericRepository: BaseRepo}
}
