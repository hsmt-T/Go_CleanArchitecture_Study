package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGorm() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
