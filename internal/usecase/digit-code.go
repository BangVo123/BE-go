package usecase

import (
	"context"
	"project/internal/models"
)

type CodeUseCase interface {
	GetCode(ctx context.Context, email, codeType string) (*models.DigitCode, error)
	CreateCode(ctx context.Context, payload *models.DigitCode) error
	DeleteCode(ctx context.Context, filter map[string]any) error
}
