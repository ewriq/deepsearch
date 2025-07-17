package handler

import (
	"deepsearch/database"
	"github.com/gofiber/fiber/v3"
)

func Search(c fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Token is required",
		})
	}

	result, err := database.RunSearch(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if result == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "no_content",
			"message": "No results found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}
