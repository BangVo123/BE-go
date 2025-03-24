package usecase

import (
	"context"
	"project/internal/models"
	"project/internal/presenter"
)

type UserCase interface {
	Login(ctx context.Context, creds presenter.LoginReq) (*models.User, error)
	SignUp(ctx context.Context, creds presenter.RegisterReq) (*models.User, error)
	CheckUserExist(ctx context.Context, filter map[string]any) (*models.User, error)
}
