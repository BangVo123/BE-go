package services

import (
	"context"
	"errors"
	"project/internal/models"
	"project/internal/repositories"
	"project/internal/usecase"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type VideoService struct {
	VideoRepo *repositories.VideoRepo
}

func NewVideoService(videoRepo *repositories.VideoRepo) usecase.VideoUseCase {
	return &VideoService{VideoRepo: videoRepo}
}

func (vs *VideoService) GetVideos(ctx context.Context, pageString, limitString string) (*[]models.Video, error) {
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	pagination := map[string]int64{
		"page":  page,
		"limit": limit,
	}
	return vs.VideoRepo.GetAll(ctx, pagination)
}

func (vs *VideoService) GetVideosWithFilter(ctx context.Context, pageString, limitString string, filter map[string]any) (*[]models.Video, error) {
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	if page == 0 {
		return vs.VideoRepo.GetMany(ctx, filter, nil)
	} else {
		limit, err := strconv.ParseInt(limitString, 10, 64)
		if err != nil {
			return nil, errors.New("Something went wrong")
		}

		pagination := map[string]int64{
			"page":  page,
			"limit": limit,
		}
		return vs.VideoRepo.GetMany(ctx, filter, pagination)
	}
}

func (vs *VideoService) AddVideo(ctx context.Context, payload models.Video) error {

	return vs.VideoRepo.Create(ctx, &payload)
}

type VideoWithOwnerInfoService struct {
	VideoWithOwnerInfoRepo *repositories.VideoWithOwnerInfoRepo
}

func NewVideoWithOwnerInfoService(VideoWithOwnerInfoRepo *repositories.VideoWithOwnerInfoRepo) usecase.VideoWithOwnerInfoUseCase {
	return &VideoWithOwnerInfoService{VideoWithOwnerInfoRepo: VideoWithOwnerInfoRepo}
}

func (vs *VideoWithOwnerInfoService) GetVideosWithOwnerInfo(ctx context.Context, pageString, limitString string) (*[]models.VideoWithOwnerInfo, error) {
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	pagination := map[string]int64{
		"page":  page,
		"limit": limit,
	}

	pipelineValue := bson.M{
		"from":         "users",
		"localField":   "belong_to",
		"foreignField": "_id",
		"as":           "belong_to",
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

	return vs.VideoWithOwnerInfoRepo.GetWithPopulation(ctx, pagination, pipelineValue, nil, "$belong_to")
}
