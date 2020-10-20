package middleware

import (
	"os"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

const jwtSecret = "qwerty123456"

func AuthRequired() func(ctx *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
	})
}
