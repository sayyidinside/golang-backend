package routers

import (
	"golang/backend/routers/apis"
	"golang/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRouter(app *fiber.App) {
	router := app.Group("")

	// Apply global middlewares here
	router.Use(cors.New(cors.Config{
		AllowOrigins:  utils.GetENVWithDefault("ALLOW_ORIGINS", "*"),
		AllowMethods:  utils.GetENVWithDefault("ALLOW_METHODS", "GET,POST,HEAD"),
		AllowHeaders:  "*",
		ExposeHeaders: "Content-Length",
	}))

	// Route list
	apis.InitAPIsRoutes(router)
}
