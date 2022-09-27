package routes

import (
	ctr "fiber-backend/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", ctr.Home)
}
