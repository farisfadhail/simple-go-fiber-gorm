package main

import (
	"go-fiber-gorm/database/migrations"
	"go-fiber-gorm/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE & MIGRATION
	migrations.RunMigration()

	app := fiber.New()
	
	routes.RouteInit(app)

	log.Fatal(app.Listen("localhost:3000"))
}