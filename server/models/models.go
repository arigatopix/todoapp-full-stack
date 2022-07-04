package models

import (
	"server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID int `gorm:"primaryKey" json:"id"`
}

func ConnectDB() *gorm.DB {
	env := config.LoadENV()

	dsn := "postgresql://" + env.POSTGRES_USER + ":" + env.POSTGRES_PASSWORD + "@" + env.POSTGRES_HOST + ":" + env.POSTGRES_PORT + "/" + env.POSTGRES_DB

	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	return db
}
