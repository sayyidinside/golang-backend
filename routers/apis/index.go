package apis

import (
	v1 "golang/backend/routers/apis/v1"

	"github.com/gofiber/fiber/v2"
)

func InitAPIsRoutes(r fiber.Router) {
	apis := r.Group("/api")

	v1.InitV1Routes(apis)
}
