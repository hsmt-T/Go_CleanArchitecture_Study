package postgres

import (
	"go_cleanArchitecture_study/domain"
	"go_cleanArchitecture_study/adapter/repository/model"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) Create(user domain.User) (domain.User, error) {

	m := model.UserModel{
		ID: string(user.ID()),
		Name: user.Name(),
		Email: user.Email(),
		CreatedAt: user.CreatedAt(),

	}

	if err := r.db.Create(&m).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserPostgres) FindByID(id domain.UserID) (domain.User, error) {
	
	var m model.UserModel

	if err := r.db.First(&m, "id = ?", string(id)).Error; err != nil {
		return domain.User{}, err
	}

	user := domain.NewUser(
		domain.UserID(m.ID),
		m.Name,
		m.Email,
		m.CreatedAt,
	)
	return user, nil
}