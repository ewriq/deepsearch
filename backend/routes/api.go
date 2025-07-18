package routes

import (
	"deepsearch/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

)

func Api(app fiber.Router) {
	app.Get("/", handler.Home)
	app.Get("/search/:token", handler.Search)
	app.Get("/ws/search", websocket.New(handler.SearchWebSocket))
}
