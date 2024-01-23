package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm/models"
)

var DB *gorm.DB

func Init() {

	var err error
	// connect to database
	dsn := "host=castor.db.elephantsql.com user=beghrxzo password=toYiNr2v9TSF4aUCaeh__hRLLyquHuhc dbname=beghrxzo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// migrate the user tables
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("cannot migrate")
	}

	err = DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic("cannot migrate")
	}
}
