package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Compress(c fiber.Ctx) error {
	compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
	return c.Next()
}

var Cors = cors.New(cors.Config{
	AllowOrigins:     []string{"http://localhost:5173"}, 
	AllowHeaders		:     	[]string{"Origin, Content-Type, Accept"},

})

func Error(c fiber.Ctx) error {
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": "500: Internal server error",
	})

	return nil
}

var Logger = logger.New(logger.Config{
    Format:     "${time} | ${pid} | ${latency} | ${status} - ${method} ${path} | ${ip} \n",
    TimeFormat: "02.01.2006 15:04:05",
})

func NotFound(c fiber.Ctx) error {
	c.Status(404).JSON(fiber.Map{
		"code":    404,
		"message": "Not Found",
	})

	return nil
}

var RateLimit = limiter.New(limiter.Config{
	Max:        1000,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c fiber.Ctx) error {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Rate limit dedected.",
		})
	},
})

func Security(c fiber.Ctx) error {
	c.Set("Content-Security-Policy", "default-src * 'self' data: blob: 'unsafe-inline' 'unsafe-eval'");
	
	return c.Next()
}
