package database

import (
	"fmt"
	"os"

	. "github.com/danudenny/golang-gorm-fiber-postgres/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%h port=%p user=%u dbname=%n password=%p", db_host, db_port, db_user, db_name, db_pass),
	)

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&User{})
	defer db.Close()
}
