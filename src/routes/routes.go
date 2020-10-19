package routes

import (
	"github.com/danudenny/golang-gorm-fiber-postgres/src/modules/user"
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/user", user.GetAllUser)
	return
}
