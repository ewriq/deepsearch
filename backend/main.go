package main

import (
	"deepsearch/middleware"
	"deepsearch/routes"
	"deepsearch/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	Config := utils.LoadConfig("./config/server.ini")
	app := fiber.New()

	app.Use(middleware.Cors)
	app.Use(middleware.Logger)
	app.Use(middleware.Compress)
	app.Use(middleware.Security)
	app.Use(middleware.RateLimit)
	app.Use(recover.New())

	service := app.Group("/")

	routes.Api(service)

	app.Use(middleware.NotFound)
	err := app.Listen(Config.Port)
	if err != nil {
		panic(err)
	}
}
