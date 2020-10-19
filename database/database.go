package database

import (
	"fmt"

	. "github.com/danudenny/golang-gorm-fiber-postgres/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=boilerplate_golang password=postgres")

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&User{})
	defer db.Close()
}
