package routers

import (
	"golang/backend/routers/apis"
	"golang/backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	router.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
	}))

	// Route list
	apis.InitAPIsRoutes(router)
}
