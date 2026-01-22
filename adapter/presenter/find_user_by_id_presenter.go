package presenter

import (
	"go_cleanArchitecture_study/domain"
	"go_cleanArchitecture_study/usecase"
)

type FindUserByIDPresenter struct {}

func NewFindUserPresenter() *FindUserByIDPresenter {
	return &FindUserByIDPresenter{}
}

func (p *FindUserByIDPresenter) Output(user domain.User) usecase.FindUserByIDOutput {
	return usecase.FindUserByIDOutput{
		ID:        string(user.ID()),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
	}
}