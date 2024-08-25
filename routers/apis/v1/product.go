package v1

import (
	"golang/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitProductRoutes(r fiber.Router) {
	productRoutes := r.Group("/products")

	productRoutes.Get("", controllers.GetProducts)
	productRoutes.Post("", controllers.AddProduct)
}
