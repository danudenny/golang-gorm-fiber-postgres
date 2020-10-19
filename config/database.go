package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=boilerplate_golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db == nil || err != nil {
		panic("Failed to connect to Database")
	}

	fmt.Println("Database connection successfully created")
}
