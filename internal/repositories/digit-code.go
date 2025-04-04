package repositories

import (
	"project/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CodeRepo struct {
	GenericRepository[models.DigitCode]
}

func NewCodeRepo(db *mongo.Database, collectionName string) *CodeRepo {
	BaseRepo := NewBaseRepo[models.DigitCode](db, collectionName)
	return &CodeRepo{GenericRepository: BaseRepo}
}
