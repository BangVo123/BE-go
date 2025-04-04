package services

import (
	"context"
	"project/internal/models"
	"project/internal/repositories"
	"project/internal/usecase"
)

type CodeServices struct {
	CodeRepo *repositories.CodeRepo
}

func NewCodeService(CodeRepo *repositories.CodeRepo) usecase.CodeUseCase {
	return &CodeServices{CodeRepo: CodeRepo}
}

func (cs *CodeServices) GetCode(ctx context.Context, email, codeType string) (*models.DigitCode, error) {
	payload := map[string]any{"email": email, "type": codeType}
	return cs.CodeRepo.GetByFilter(ctx, payload)
}
func (cs *CodeServices) CreateCode(ctx context.Context, payload *models.DigitCode) error {
	return cs.CodeRepo.Create(ctx, payload)
}
func (cs *CodeServices) DeleteCode(ctx context.Context, filter map[string]any) error {
	return cs.CodeRepo.DeleteByFilter(ctx, filter)
}
