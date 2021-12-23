package router

import (
	"github.com/gofiber/fiber/v2"
)

func InstallRouter(app *fiber.App) {
	NewHttpRouter().InstallRouter(app)
	NewApiRouter().InstallRouter(app)
}
