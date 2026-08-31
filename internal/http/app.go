package httpapp

import (
	"rida-api/internal/config"

	"github.com/gofiber/fiber/v3"
)

func NewApp(cfg config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	registerHealthRoutes(app)
	registerSwagger(app, cfg)

	return app
}
