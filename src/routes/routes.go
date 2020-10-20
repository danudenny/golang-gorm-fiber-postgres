package routes

import (
	. "github.com/danudenny/golang-gorm-fiber-postgres/src/middleware"
	. "github.com/danudenny/golang-gorm-fiber-postgres/src/modules/auth"
	. "github.com/danudenny/golang-gorm-fiber-postgres/src/modules/user"
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Post("api/v1/login", Login)

	app.Get("/api/v1/user", AuthRequired(), GetAllUser)
}
