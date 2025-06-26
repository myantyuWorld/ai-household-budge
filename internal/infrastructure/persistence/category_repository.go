package persistence

import (
	"ai-household-budge/internal/domain/model"
	"ai-household-budge/internal/domain/repository"
	"context"
	"errors"
	"sync"
)

type CategoryRepositoryImpl struct {
	categories map[string]*model.Category
	mutex      sync.RWMutex
}

func NewCategoryRepository() repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		categories: make(map[string]*model.Category),
	}
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, category *model.Category) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.categories[category.ID]; exists {
		return errors.New("category already exists")
	}

	// ディープコピーを作成
	categoryCopy := *category
	r.categories[category.ID] = &categoryCopy

	return nil
}

func (r *CategoryRepositoryImpl) GetByID(ctx context.Context, id string) (*model.Category, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	category, exists := r.categories[id]
	if !exists {
		return nil, errors.New("category not found")
	}

	// ディープコピーを返す
	categoryCopy := *category
	return &categoryCopy, nil
}

func (r *CategoryRepositoryImpl) GetAll(ctx context.Context) ([]*model.Category, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	categories := make([]*model.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categoryCopy := *category
		categories = append(categories, &categoryCopy)
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, category *model.Category) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.categories[category.ID]; !exists {
		return errors.New("category not found")
	}

	// ディープコピーを作成
	categoryCopy := *category
	r.categories[category.ID] = &categoryCopy

	return nil
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.categories[id]; !exists {
		return errors.New("category not found")
	}

	delete(r.categories, id)
	return nil
}
