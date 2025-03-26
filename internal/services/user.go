package services

import (
	"context"
	"errors"
	"project/internal/models"
	"project/internal/presenter"
	"project/internal/repositories"
	"project/internal/usecase"
)

type AuthService struct {
	UserRepo *repositories.UserRepo
}

func NewAuthService(userRepo *repositories.UserRepo) usecase.UserCase {
	return &AuthService{UserRepo: userRepo}
}

func (as *AuthService) Login(ctx context.Context, creds presenter.LoginReq) (*models.User, error) {
	Username := creds.Username
	Password := creds.Password

	FoundUser, err := as.UserRepo.GetByField(ctx, "email", Username)
	if err != nil {
		return nil, err
	}

	isTrue := FoundUser.ComparePassword(Password)
	if !isTrue {
		return nil, errors.New("Username or password does not match")
	}

	return FoundUser, nil
}

func (as *AuthService) SignUp(ctx context.Context, creds presenter.RegisterReq) (*models.User, error) {
	//code for test
	Username := creds.Username
	Password := creds.Password

	FoundUser, err := as.UserRepo.GetByField(ctx, "email", Username)
	if err != nil {
		return nil, err
	}

	isTrue := FoundUser.ComparePassword(Password)
	if !isTrue {
		return nil, errors.New("Username or password does not match")
	}

	return FoundUser, nil
}

func (as *AuthService) GetUserExist(ctx context.Context, filter map[string]any) (*models.User, error) {
	FoundUser, err := as.UserRepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return FoundUser, nil
}

func (as *AuthService) CreateUser(ctx context.Context, payload *models.User) (*models.User, error) {
	return as.CreateUser(ctx, payload)
}

func (as *AuthService) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return as.UserRepo.GetById(ctx, userId)
}
