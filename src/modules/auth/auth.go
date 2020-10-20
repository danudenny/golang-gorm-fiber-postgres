package auth

import (
	"os"
	"time"

	login "github.com/danudenny/golang-gorm-fiber-postgres/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

const jwtSecret = "qwerty123456"

func Login(ctx *fiber.Ctx) {
	var bodyRequest login.LoginDTO
	err := ctx.BodyParser(&bodyRequest)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
		return
	}

	if bodyRequest.Username != "danudenny" || bodyRequest.Password != "admin" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24)

	s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
	})

}
