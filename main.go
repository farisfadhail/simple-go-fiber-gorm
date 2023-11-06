package main

import (
	"go-fiber-gorm/database/migrations"
	"go-fiber-gorm/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE & MIGRATION
	migrations.RunMigration()

	app := fiber.New()
	
	routes.RouteInit(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Println("Failed to listen go fiber server")
		os.Exit(1)
	}
}