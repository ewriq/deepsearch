package api

import (
	"deepsearch/middleware"
	"deepsearch/routes"
	"deepsearch/utils"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func Start() error {
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
fmt.Println(Config.Port, " is the port the server is running on.")
	err := app.Listen(Config.Port)
	if err != nil {
		panic(err)
	}

	return err
}
