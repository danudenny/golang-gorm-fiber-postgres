package user

import (
	"github.com/danudenny/golang-gorm-fiber-postgres/src/middleware"
	fiber "github.com/gofiber/fiber"
)

func GetAllUser(c *fiber.Ctx) {
	middleware.AuthRequired()
	c.Send("Nice")
}
