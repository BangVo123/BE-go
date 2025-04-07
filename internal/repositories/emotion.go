package repositories

import (
	"project/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepo struct {
	GenericRepository[models.Emotion]
}

func NewLikeRepo(db *mongo.Database, collectionName string) *LikeRepo {
	baseRepo := NewBaseRepo[models.Emotion](db, collectionName)
	return &LikeRepo{GenericRepository: baseRepo}
}

type LoveRepo struct {
	GenericRepository[models.Emotion]
}

func NewLoveRepo(db *mongo.Database, collectionName string) *LoveRepo {
	baseRepo := NewBaseRepo[models.Emotion](db, collectionName)
	return &LoveRepo{GenericRepository: baseRepo}
}
