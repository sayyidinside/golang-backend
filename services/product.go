package services

import (
	"golang/backend/database"
	"golang/backend/dtos"
	"golang/backend/helpers"
	"golang/backend/models"
	"golang/backend/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func FetchProducts(query *dtos.QueryDTO, url *string) *dtos.Response {
	var products []models.Product
	var data interface{}
	var totalRows int64

	dbQuery := database.DBConn.Model(&models.Product{}).
		Joins("JOIN product_categories ON product_categories.id = products.category_id").
		Preload("Category", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		})

	// Apply pagination
	dbQuery = dbQuery.Scopes(helpers.Paginate(query))

	// Apply order
	{
		var allowedFields = map[string]string{
			"created": "products.created_at",
		}

		// ! hardcode the order value
		query.OrderBy = "created"
		query.Order = "desc"

		dbQuery = dbQuery.Scopes(helpers.Order(query, allowedFields))
	}

	// Apply Search conditionally
	if query.Search != "" {
		search := "%" + query.Search + "%"
		dbQuery = dbQuery.Where("products.name LIKE ?", search).Or("products.description LIKE ?", search)
	}

	// Apply fillter conditionally
	if query.Filter != "" {
		var allowedFields = map[string]string{
			"category": "products.category_id",
		}

		// ! hardcode the order value
		query.FilterBy = "category"

		dbQuery = dbQuery.Scopes(helpers.Filter(query, allowedFields))
	}

	// Retrieving data
	if err := dbQuery.Find(&products).Error; err != nil {
		return &dtos.Response{
			Success: false,
			Message: "Failed retrieving data",
			Error:   err.Error(),
		}
	}

	// Counting total data
	{
		countQuery := database.DBConn.Model(&models.Product{}).Joins("JOIN product_categories ON product_categories.id = products.category_id")

		if query.Search != "" {
			search := "%" + query.Search + "%"
			countQuery = countQuery.Where("products.name LIKE ?", search).Or("products.description LIKE ?", search)
		}

		if query.Filter != "" {
			countQuery = countQuery.Where("products.category_id = ?", query.Filter)
		}

		countQuery.Count(&totalRows)
	}

	productDtos := dtos.ToProductDTOs(&products)
	data = productDtos

	data = helpers.GeneratePaginatedQuery(query, url, totalRows, utils.ToInterfaceSlice(data))

	return &dtos.Response{
		Success: true,
		Message: "Product succesfully retrieve",
		Data:    data,
	}
}
