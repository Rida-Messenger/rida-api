package httpapp

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func registerHealthRoutes(app *fiber.App) {
	app.Get(
		healthcheck.LivenessEndpoint,
		healthcheck.New(),
	)
}
