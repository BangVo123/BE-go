package services

import (
	"context"
	"project/internal/models"
	"project/internal/repositories"
	"project/internal/usecase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FollowService struct {
	FollowRepo *repositories.FollowerRepo
}

func NewFollowService(FollowRepo *repositories.FollowerRepo) usecase.FollowUseCase {
	return &FollowService{FollowRepo: FollowRepo}
}

func (fl *FollowService) GetAll(ctx context.Context, userId string) (*[]models.Follower, error) {
	return fl.FollowRepo.GetAll(ctx, nil)
}

type FollowerInfoService struct {
	FollowerInfoRepo *repositories.FollowerInfoRepo
}

func NewFollowerInfoService(FollowerInfoRepo *repositories.FollowerInfoRepo) usecase.FollowerInfoUseCase {
	return &FollowerInfoService{FollowerInfoRepo: FollowerInfoRepo}
}

func (fler *FollowerInfoService) GetFollower(ctx context.Context, userId string) (*[]models.FollowerInfo, error) {
	ObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	pipelineVal := bson.M{
		"from":         "users",
		"localField":   "user_id",
		"foreignField": "_id",
		"as":           "user_id",
		"pipeline": bson.A{
			bson.D{
				{Key: "$project", Value: bson.D{
					{Key: "_id", Value: 1},
					{Key: "full_name", Value: 1},
					{Key: "nick_name", Value: 1},
					{Key: "avatar", Value: 1},
					{Key: "bio", Value: 1},
					{Key: "tick", Value: 1},
					{Key: "followers_count", Value: 1},
					{Key: "likes_count", Value: 1},
				}},
			},
		},
	}

	filter := map[string]any{
		"following_id": ObjectId,
	}

	return fler.FollowerInfoRepo.GetWithPopulation(ctx, nil, pipelineVal, filter, "$user_id")
}

type FollowingInfoService struct {
	FollowingInfoRepo *repositories.FollowingInfoRepo
}

func NewFollowingInfoService(FollowingInfoRepo *repositories.FollowingInfoRepo) usecase.FollowingInfoUseCase {
	return &FollowingInfoService{FollowingInfoRepo: FollowingInfoRepo}
}

func (fling *FollowingInfoService) GetFollowing(ctx context.Context, userId string) (*[]models.FollowingInfo, error) {
	ObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	pipelineVal := bson.M{
		"from":         "users",
		"localField":   "following_id",
		"foreignField": "_id",
		"as":           "following_id",
		"pipeline": bson.A{
			bson.D{
				{Key: "$project", Value: bson.D{
					{Key: "_id", Value: 1},
					{Key: "full_name", Value: 1},
					{Key: "nick_name", Value: 1},
					{Key: "email", Value: 1},
					{Key: "phone_nums", Value: 1},
					{Key: "password", Value: 1},
					{Key: "avatar", Value: 1},
					{Key: "bio", Value: 1},
					{Key: "account_id", Value: 1},
					{Key: "provider", Value: 1},
					{Key: "tick", Value: 1},
					{Key: "followings_count", Value: 1},
					{Key: "followers_count", Value: 1},
					{Key: "likes_count", Value: 1},
					{Key: "website_URL", Value: 1},
					{Key: "facebook_URL", Value: 1},
					{Key: "youtube_URL", Value: 1},
					{Key: "twitter_URL", Value: 1},
					{Key: "instagram_URL", Value: 1},
				}},
			},
		},
	}

	filter := map[string]any{
		"user_id": ObjectId,
	}

	return fling.FollowingInfoRepo.GetWithPopulation(ctx, nil, pipelineVal, filter, "$following_id")
}
