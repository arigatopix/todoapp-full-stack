package models

import (
	config "server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	// CreatedOn  int `json:"created_on"`
	// ModifiedOn int `json:"modified_on"`
	// DeletedOn  int `json:"deleted_on"`
}

var db *gorm.DB

func ConnectDB() {
	env := config.LoadENV()

	var err error

	dsn := "host=" + env.POSTGRES_HOST + " user=" + env.POSTGRES_USER + " password=" + env.POSTGRES_PASSWORD + " dbname=" + env.POSTGRES_DB + " port=" + env.POSTGRES_PORT + " sslmode=disable TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Task{})
}
