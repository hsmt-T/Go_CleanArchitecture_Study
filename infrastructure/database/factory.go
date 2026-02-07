package database

import (
	"go_cleanArchitecture_study/adapter/repository"
	"go_cleanArchitecture_study/adapter/repository/postgres"
	"go_cleanArchitecture_study/domain"
	"os"
)

func NewUserRepository() domain.UserRepository {

	driver := os.Getenv("DB_DRIVER")

	switch driver {
		
	case "postgres":
		db := NewPostgresGorm()
		return postgres.NewUserPostgres(db)
	

	default:
		return repository.NewUserMemoryRepository()
	}
}