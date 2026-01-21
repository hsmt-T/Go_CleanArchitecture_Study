package presenter

import (
	"go_cleanArchitecture_study/domain"
	"go_cleanArchitecture_study/usecase"
)

type CreateUserPresenter struct{}

func NewCreateUserPresenter() *CreateUserPresenter {
	return &CreateUserPresenter{}
}

func (p *CreateUserPresenter) Output(user domain.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:        string(user.ID()),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
	}
}