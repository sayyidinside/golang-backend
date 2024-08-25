package controllers

import (
	"golang/backend/dtos"
	"golang/backend/helpers"
	"golang/backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

func GetProducts(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + c.OriginalURL()

	sanitizer := bluemonday.UGCPolicy()
	response := services.FetchProducts(
		&dtos.QueryDTO{
			Page:     sanitizer.Sanitize(c.Query("page")),
			Limit:    sanitizer.Sanitize(c.Query("limit")),
			Search:   sanitizer.Sanitize(c.Query("search")),
			FilterBy: sanitizer.Sanitize(c.Query("filter_by")),
			Filter:   sanitizer.Sanitize(c.Query("filter")),
			OrderBy:  sanitizer.Sanitize(c.Query("order_by")),
			Order:    sanitizer.Sanitize(c.Query("order")),
		},
		&baseURL,
	)

	if response.Success {
		c.Status(fiber.StatusOK)
	} else {
		c.Status(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}

func AddProduct(c *fiber.Ctx) error {
	var input dtos.InputProductDTO

	// Sanitize input fields
	sanitizer := bluemonday.UGCPolicy()
	input.Name = sanitizer.Sanitize(input.Name)
	input.Description = sanitizer.Sanitize(input.Description)
	input.CategoryID = sanitizer.Sanitize(input.CategoryID)

	if err := c.BodyParser(&input); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(dtos.Response{
				Success: false,
				Message: "Failed parsing user input",
				Data:    nil,
				Error:   err.Error(),
			})
	}

	if res, err := helpers.ValidateInput(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	response := services.InsertProduct(&input)

	if response.Success {
		c.Status(fiber.StatusCreated)
	} else if response.Message == "Invalid data" {
		c.Status(fiber.StatusBadRequest)
	} else {
		c.Status(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}
