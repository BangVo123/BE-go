package usecase

import (
	"context"
	"project/internal/models"
)

type VideoUseCase interface {
	GetVideos(ctx context.Context, page, limit string) (*[]models.Video, error)
}
