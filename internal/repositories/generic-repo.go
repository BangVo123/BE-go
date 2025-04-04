package repositories

import "context"

type GenericRepository[T any] interface {
	GetCollection() any
	GetById(ctx context.Context, id string) (*T, error)
	GetByField(ctx context.Context, field string, value any) (*T, error)
	GetByFilter(ctx context.Context, filter map[string]any) (*T, error)
	Create(ctx context.Context, payload *T) error
	Update(ctx context.Context, id string, payload map[string]any) (*T, error)
	Delete(ctx context.Context, id string) error
	DeleteByFilter(ctx context.Context, filter map[string]any) error
	GetMany(ctx context.Context, filter map[string]any, pagination map[string]int64) (*[]T, error)
	GetAll(ctx context.Context, pagination map[string]int64) (*[]T, error)
	GetWithPopulation(ctx context.Context, pagination map[string]int64, pipelineValue map[string]any, unwindAttr string) (*[]T, error)
}
