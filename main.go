package main

import (
	"fmt"

	"github.com/danudenny/golang-gorm-fiber-postgres/src/modules/user"
	fiber "github.com/gofiber/fiber"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

func initMain(c *fiber.Ctx) {
	c.Send("First Init")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/user", user.GetAllUser)
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=boilerplate_golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db == nil || err != nil {
		panic("Failed to connect to Database")
	}

	fmt.Println("Database connection successfully created")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
}
