package config

import (
	"github.com/gofiber/fiber/v2"
)

// NewApp membuat instance aplikasi Fiber.
func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Gym Management System",
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Gym Management System API is running",
		})
	})

	return app
}
