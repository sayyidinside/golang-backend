package main

import (
	"golang/backend/database"
	"golang/backend/routers"
	"golang/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	utils.LoadENV()
	database.ConnectDb()

	routers.SetupRouter(app)

	app.Listen(utils.GetENVWithDefault("port", ":3000"))
}
