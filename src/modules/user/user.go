package user

import (
	fiber "github.com/gofiber/fiber"
)

func GetAllUser(c *fiber.Ctx) {
	c.Send("Nice")
}
