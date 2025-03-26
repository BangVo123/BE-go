package usecase

import (
	"context"
	"project/internal/models"
	"project/internal/presenter"
)

type UserCase interface {
	Login(ctx context.Context, creds presenter.LoginReq) (*models.User, error)
	SignUp(ctx context.Context, creds presenter.RegisterReq) (*models.User, error)
	GetUserExist(ctx context.Context, filter map[string]any) (*models.User, error)
	CreateUser(ctx context.Context, payload *models.User) (*models.User, error)
	GetUserById(ctx context.Context, userId string) (*models.User, error)
}
