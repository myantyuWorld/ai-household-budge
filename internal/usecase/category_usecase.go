package usecase

import (
	"ai-household-budge/internal/domain/model"
	"ai-household-budge/internal/domain/repository"
	"context"
	"errors"
)

type CategoryUseCase struct {
	repo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		repo: repo,
	}
}

func (uc *CategoryUseCase) Create(ctx context.Context, name, description, color string) (*model.Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	category := model.NewCategory(name, description, color)

	if err := uc.repo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *CategoryUseCase) GetByID(ctx context.Context, id string) (*model.Category, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return uc.repo.GetByID(ctx, id)
}

func (uc *CategoryUseCase) GetAll(ctx context.Context) ([]*model.Category, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *CategoryUseCase) Update(ctx context.Context, id, name, description, color string) (*model.Category, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	category, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		category.UpdateName(name)
	}
	if description != "" {
		category.UpdateDescription(description)
	}
	if color != "" {
		category.UpdateColor(color)
	}

	if err := uc.repo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *CategoryUseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return uc.repo.Delete(ctx, id)
}
