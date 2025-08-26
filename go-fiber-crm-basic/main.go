package main

import (
	"fmt"

	"github.com/eswar/go-fiber-crm-basic/database"
	"github.com/eswar/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()

	// Connect to database
	database.ConnectDB()
	fmt.Println("Database connection established")

	// Auto-migrate models
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")

	// Routes
	setupRoutes(app)

	// Start server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
