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
