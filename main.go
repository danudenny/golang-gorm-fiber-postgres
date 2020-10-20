package main

import (
	"github.com/danudenny/golang-gorm-fiber-postgres/database"
	"github.com/danudenny/golang-gorm-fiber-postgres/src/routes"
	fiber "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func initMain(c *fiber.Ctx) {
	c.Send("First Init")
}

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	database.Connect()
	routes.SetupRoutes(app)
	app.Listen(3000)
}
