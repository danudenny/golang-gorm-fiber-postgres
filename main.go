package main

import (
	"github.com/danudenny/golang-boilerplate/src/modules/user"
	fiber "github.com/gofiber/fiber"
)

func initMain(c *fiber.Ctx) {
	c.Send("First Init")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/user", user.GetAllUser)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
