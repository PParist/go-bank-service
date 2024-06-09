package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func ApiMiddleware(c *fiber.Ctx) error {
	// Logging
	log.Printf("API Request: %s %s", c.Method(), c.OriginalURL())

	// CORS settings
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Authentication check (example: simple token-based auth)
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Proceed to the next middleware or handler
	return c.Next()
}

// RateLimiter middleware for API
func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,             // maximum number of requests per duration
		Expiration: 1 * time.Minute, // reset count every minute
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too Many Requests",
			})
		},
	})
}

// Compression middleware for API responses
func Compression() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // compression level
	})
}
