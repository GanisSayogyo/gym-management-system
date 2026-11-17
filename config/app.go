package config

import (
	"github.com/gofiber/fiber/v2"
)

// NewApp membuat instance aplikasi Fiber.
func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Gym Management System",
	})

	return app
}
