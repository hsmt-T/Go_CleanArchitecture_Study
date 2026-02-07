package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGorm() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
