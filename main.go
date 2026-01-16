package main

import (
	"log"
	"votepeace/database"
	"votepeace/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Connect to Database & AutoMigrate
	database.Connect()

	// 2. Initialize Fiber App
	app := fiber.New()

	// 3. Middlewares
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173", // Frontend URL
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	// 4. Routes
	routes.Setup(app)

	// 5. Start Server
	log.Fatal(app.Listen(":3000"))
}
