package routers

import (
	"golang/backend/routers/apis"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	router := app.Group("")
	apis.InitAPIsRoutes(router)
	// InitV1Route
	// Apply global middlewares here
}
