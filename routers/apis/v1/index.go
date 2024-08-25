package v1

import (
	"github.com/gofiber/fiber/v2"
)

func InitV1Routes(r fiber.Router) {
	v1Routes := r.Group("/v1")

	InitProductRoutes(v1Routes)
}
