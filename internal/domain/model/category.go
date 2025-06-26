package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCategory(name, description, color string) *Category {
	return &Category{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Color:       color,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (c *Category) UpdateName(name string) {
	c.Name = name
	c.UpdatedAt = time.Now()
}

func (c *Category) UpdateDescription(description string) {
	c.Description = description
	c.UpdatedAt = time.Now()
}

func (c *Category) UpdateColor(color string) {
	c.Color = color
	c.UpdatedAt = time.Now()
}
