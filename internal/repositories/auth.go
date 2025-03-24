package repositories

import (
	"context"
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

func (ur *UserRepo) GetUserByFilter(ctx context.Context, filter map[string]any) (*models.User, error) {
	return ur.GetByFilter(ctx, filter)
}
