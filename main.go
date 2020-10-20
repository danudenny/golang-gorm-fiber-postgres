package main

import (
	"log"

	"github.com/danudenny/golang-gorm-fiber-postgres/database"
	"github.com/danudenny/golang-gorm-fiber-postgres/src/routes"
	fiber "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"
)

func initMain(c *fiber.Ctx) {
	c.Send("First Init")
}

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	app.Use(middleware.Logger())
	database.Connect()
	routes.SetupRoutes(app)
	app.Listen(3000)
}
