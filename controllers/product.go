package controllers

import "github.com/gofiber/fiber/v2"

func GetProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    "test",
	})
}
