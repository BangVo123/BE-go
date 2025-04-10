package usecase

import (
	"context"
	"project/internal/models"
)

type VideoUseCase interface {
	GetVideos(ctx context.Context, pageString, limitString string) (*[]models.Video, error)
	GetVideosWithFilter(ctx context.Context, pageString, limitString string, filter map[string]any) (*[]models.Video, error)
	AddVideo(ctx context.Context, payload models.Video) error
}

type VideoWithOwnerInfoUseCase interface {
	GetVideosWithOwnerInfo(ctx context.Context, pageString, limitString string) (*[]models.VideoWithOwnerInfo, error)
}
