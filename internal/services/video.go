package services

import (
	"context"
	"errors"
	"project/internal/models"
	"project/internal/repositories"
	"project/internal/usecase"
	"strconv"
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
