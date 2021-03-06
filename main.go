package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gofiber-gorm-mysql/database"
	"gofiber-gorm-mysql/router"
)

func main() {
	// Start new fiber app
	app := fiber.New()

	// Use logger middleware
	app.Use(logger.New())

	// Connect to the database
	database.ConnectDB()
	router.SetupRoutes(app)

	// Listen on port 3000
	err := app.ListenTLS(":3000", "/etc/letsencrypt/live/www.alifnuryana.software/fullchain.pem", "/etc/letsencrypt/live/www.alifnuryana.software/privkey.pem")
	if err != nil {
		return
	}

}
