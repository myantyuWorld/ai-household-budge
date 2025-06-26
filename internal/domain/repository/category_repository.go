package repository

import (
	"ai-household-budge/internal/domain/model"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *model.Category) error
	GetByID(ctx context.Context, id string) (*model.Category, error)
	GetAll(ctx context.Context) ([]*model.Category, error)
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id string) error
}
