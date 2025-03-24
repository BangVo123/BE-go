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
	UserRepo repositories.UserRepo
}

func NewAuthService(userRepo repositories.UserRepo) usecase.UserCase {
	return &AuthService{UserRepo: userRepo}
}

func (ah *AuthService) Login(ctx context.Context, creds presenter.LoginReq) (*models.User, error) {
	Username := creds.Username
	Password := creds.Password

	FoundUser, err := ah.UserRepo.GetByField(ctx, "email", Username)
	if err != nil {
		return nil, err
	}

	// v := reflect.ValueOf(FoundUser)
	// // Make sure we're dealing with a struct
	// if v.Kind() == reflect.Ptr {
	// 	v = v.Elem()
	// }
	// t := v.Type()

	// for i := 0; i < v.NumField(); i++ {
	// 	field := t.Field(i)
	// 	value := v.Field(i).Interface()
	// 	fmt.Printf("%s: %v\n", field.Name, value)
	// }

	isTrue := FoundUser.ComparePassword(Password)
	if !isTrue {
		return nil, errors.New("Username or password does not match")
	}

	return FoundUser, nil
}

func (ah *AuthService) SignUp(ctx context.Context, creds presenter.RegisterReq) (*models.User, error) {
	//code for test
	Username := creds.Username
	Password := creds.Password

	FoundUser, err := ah.UserRepo.GetByField(ctx, "email", Username)
	if err != nil {
		return nil, err
	}

	isTrue := FoundUser.ComparePassword(Password)
	if !isTrue {
		return nil, errors.New("Username or password does not match")
	}

	return FoundUser, nil
}

func (ah *AuthService) CheckUserExist(ctx context.Context, filter map[string]any) (*models.User, error) {
	FoundUser, err := ah.UserRepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return FoundUser, nil
}
