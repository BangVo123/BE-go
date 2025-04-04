package repositories

import (
	"project/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	GenericRepository[models.User]
}

func NewUserRepo(db *mongo.Database, collectionName string) *UserRepo {
	BaseRepo := NewBaseRepo[models.User](db, collectionName)
	return &UserRepo{GenericRepository: BaseRepo}
}
