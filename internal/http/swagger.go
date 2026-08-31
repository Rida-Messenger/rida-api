package httpapp

import (
	"fmt"
	"rida-api/docs"
	"rida-api/internal/config"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func registerSwagger(app *fiber.App, cfg config.Config) error {
	docs.SwaggerInfo.Title = cfg.App.Name
	docs.SwaggerInfo.Description = fmt.Sprintf("OpenAPI schema for %s", cfg.App.Name)
	docs.SwaggerInfo.Version = cfg.App.Version
	docs.SwaggerInfo.BasePath = cfg.App.RootPath

	app.Get("/docs/*", swaggo.New(swaggo.Config{
		Title: cfg.App.Name,
	}))

	return nil
}
