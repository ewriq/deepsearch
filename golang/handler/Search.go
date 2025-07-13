package handler

import (
	"deepsearch/database"
	"github.com/gofiber/fiber/v3"
)

func Search(c fiber.Ctx) error {
	token := c.Params("token")
	if token == "" { 
		return c.JSON(fiber.Map{"status": "no content", "message": "Token is required"})
	}

	aram, err := database.RunSearch(token)
	if aram != "" {
		return c.JSON(fiber.Map{"status": "success", "data": aram})
	} else {
		return c.JSON(fiber.Map{"status": "error", "message": err.Error()})

	}
	
}
