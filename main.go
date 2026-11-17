package main

import (
	"context"
	"os"

	"github.com/GanisSayogyo/gym-management-system/config"
	"github.com/GanisSayogyo/gym-management-system/database"
)

func main() {
	config.LoadEnv()

	logger := config.NewLogger()
	logger.Info("Gym Management System starting")

	ctx := context.Background()

	pool, err := database.NewPool(ctx)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	logger.Info("Database connection established successfully")

	app := config.NewApp()

	logger.Info("Fiber application initialized",
		"app_name", app.Config().AppName,
	)

	port := config.GetEnv("APP_PORT", "3000")

	logger.Info("HTTP server starting", "port", port)

	if err := app.Listen(":" + port); err != nil {
		logger.Error("HTTP server failed", "error", err)
		os.Exit(1)
	}
}
