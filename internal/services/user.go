package services

import (
	"context"
	"errors"
	"project/internal/models"
	"project/internal/presenter"
	"project/internal/repositories"
	"project/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (as *AuthService) SignUp(ctx context.Context, creds presenter.RegisterReq) (string, error) {
	FoundUser, err := as.UserRepo.GetByField(ctx, "email", creds.Email)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return "", err
	}
	if FoundUser != nil {
		return "", errors.New("User already exists")
	}

	userId := primitive.NewObjectID()
	payload := models.User{
		Id:       userId,
		Email:    creds.Email,
		Password: creds.Password,
	}
	err = payload.HashPassword()
	if err != nil {
		return "", err
	}

	err = as.UserRepo.Create(ctx, &payload)
	if err != nil {
		return "", nil
	}

	return userId.Hex(), nil
}

func (as *AuthService) Reset(ctx context.Context, email, password string) error {
	foundUser, err := as.GetUserExist(ctx, map[string]any{"email": email})
	if err != nil {
		return err
	}

	foundUser.Password = password
	err = foundUser.HashPassword()
	if err != nil {
		return err
	}

	_, err = as.UserRepo.Update(ctx, foundUser.Id.Hex(), map[string]any{"password": foundUser.Password, "provider": nil})
	return err
}

func (as *AuthService) GetUserExist(ctx context.Context, filter map[string]any) (*models.User, error) {
	FoundUser, err := as.UserRepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return FoundUser, nil
}

func (as *AuthService) CreateUser(ctx context.Context, payload *models.User) error {
	return as.UserRepo.Create(ctx, payload)
}

func (as *AuthService) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return as.UserRepo.GetById(ctx, userId)
}

func (as *AuthService) UpdateUser(ctx context.Context, userId string, payload map[string]any) error {
	_, err := as.UserRepo.Update(ctx, userId, payload)
	return err
}
