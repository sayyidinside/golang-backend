package dtos

import (
	"golang/backend/models"

	"github.com/google/uuid"
)

type (
	ProductsDTO struct {
		ID          uuid.UUID `json:"id"`
		CategoryID  uuid.UUID `json:"category_id"`
		Name        string    `json:"name"`
		Category    string    `json:"category"`
		Description string    `json:"description"`
	}

	InputProductDTO struct {
		CategoryID  string `json:"category_id" form:"category_id" validate:"required"`
		Name        string `json:"name" form:"name" validate:"required"`
		Description string `json:"description"`
	}
)

func ToProductKDTO(model models.Product) ProductsDTO {
	return ProductsDTO{
		ID:          model.ID,
		CategoryID:  model.CategoryID,
		Name:        model.Name,
		Category:    model.Category.Name,
		Description: model.Description,
	}
}

func ToProductDTOs(models []models.Product) []ProductsDTO {
	productDTOs := make([]ProductsDTO, len(models))

	for index, model := range models {
		productDTOs[index] = ToProductKDTO(model)
	}

	return productDTOs
}
