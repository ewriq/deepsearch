package handler

import (
	"deepsearch/database"
	"net/url"

	"github.com/gofiber/fiber/v3"
)

func Search(c fiber.Ctx) error {
	encodedToken := c.Params("token")
	if encodedToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Token is required",
		})
	}

	token, err := url.QueryUnescape(encodedToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid token parameter",
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
