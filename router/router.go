package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "gofiber-gorm-mysql/internal/routes/note"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())
	// Setup Note Routes
	noteRoutes.SetupNoteRouter(api)
}
