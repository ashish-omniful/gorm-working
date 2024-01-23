package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm/models"
)

var DB *gorm.DB

func main() {

	var err error
	fmt.Println("check")

	// connect to database
	dsn := "host=castor.db.elephantsql.com user=beghrxzo password=toYiNr2v9TSF4aUCaeh__hRLLyquHuhc dbname=beghrxzo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// migrate the user tables
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("cannot migrate")
	}
}
