package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter(limiterPath string) limiter.Config {
	switch limiterPath {
	case "/auth/forgot-password":
		return limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return false
			},
			Max:        10,
			Expiration: 6 * time.Hour,
			KeyGenerator: func(c *fiber.Ctx) string {
				return string(c.Body())
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(429).JSON("limited request")
			},
		}
	case "/auth/get-otp":
		return limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return false
			},
			Max:        10,
			Expiration: 1 * time.Hour,
			KeyGenerator: func(c *fiber.Ctx) string {
				return string(c.Body())
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(429).JSON("limited request")
			},
		}
	default:
		return limiter.Config{}
	}
}
