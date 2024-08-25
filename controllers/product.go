package controllers

import (
	"golang/backend/dtos"
	"golang/backend/helpers"
	"golang/backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	response := services.FetchProducts()

	if response.Success {
		c.Status(fiber.StatusOK)
	} else {
		c.Status(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}

func AddProduct(c *fiber.Ctx) error {
	var input dtos.InputProductDTO
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
