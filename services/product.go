package services

import (
	"golang/backend/database"
	"golang/backend/dtos"
	"golang/backend/helpers"
	"golang/backend/models"

	"github.com/google/uuid"
)

func InsertProduct(input *dtos.InputProductDTO) *dtos.Response {
	// check existence of category id
	var categoryCount int64
	database.DBConn.Model(&models.ProductCategory{}).Where("id = ?", input.CategoryID).Count(&categoryCount)
	if categoryCount == 0 {
		return &dtos.Response{
			Success: false,
			Message: "Invalid data",
			Error: helpers.ErrorResponse{
				Field: "category_id",
				Tag:   "not_found",
				Value: input.CategoryID,
			},
		}
	}

	product := models.Product{
		CategoryID:  uuid.MustParse(input.CategoryID),
		Name:        input.Name,
		Description: input.Description,
	}

	if err := database.DBConn.Create(&product).Error; err != nil {
		return &dtos.Response{
			Success: false,
			Message: "Failed creating data",
			Error:   err.Error(),
		}
	}

	return &dtos.Response{
		Success: true,
		Message: "Product successfully created",
	}
}

func FetchProducts() *dtos.Response {
	var response dtos.Response

	return &response
}
