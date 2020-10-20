package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger"
	"github.com/danudenny/golang-gorm-fiber-postgres/database"
	_ "github.com/danudenny/golang-gorm-fiber-postgres/docs"
	"github.com/danudenny/golang-gorm-fiber-postgres/src/routes"
	"github.com/gofiber/cors"
	fiber "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New(&fiber.Settings{
		Immutable: true,
	})

	app.Use(cors.New())

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	app.Use("/swagger", swagger.New())

	app.Use(middleware.Logger())
	database.Connect()
	routes.SetupRoutes(app)
	app.Listen(3000)
}
