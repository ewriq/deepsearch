package routes

import (
	"deepsearch/handler"

	"github.com/gofiber/fiber/v3"
)

func Api(app fiber.Router) {
	app.Get("/", handler.Home)
	app.Get("/search/:token", handler.Search)

}
