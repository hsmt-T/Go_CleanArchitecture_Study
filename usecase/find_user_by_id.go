package usecase

import (
	"go_cleanArchitecture_study/domain"
	"time"
)

type FindUserByIDInput struct {
	ID string
}

type FindUserByIDOutput struct {
	ID			string
	Name		string
	Email		string
	CreatedAt	time.Time
}

type FindUserByIDUseCase interface{
	Execute(input FindUserByIDInput) (FindUserByIDOutput, error)
}

type FindUserByIDPresenter interface {
	Output(domain.User) FindUserByIDOutput
}

// domain に依存する/repository を使う が　抽象に依存
type findUserByIDInteractor struct {
	userRepo domain.UserRepository
	presenter FindUserByIDPresenter
}
//コンストラクタ

func NewFindUserByIDInteractor(repo domain.UserRepository, presenter FindUserByIDPresenter) FindUserByIDUseCase {
	return &findUserByIDInteractor{
		userRepo: repo,
		presenter: presenter,
	}
}

func (i *findUserByIDInteractor) Execute(input FindUserByIDInput) (FindUserByIDOutput, error) {

	user, err := i.userRepo.FindByID(domain.UserID(input.ID))

	if err != nil {
		return FindUserByIDOutput{}, err
	}

	return i.presenter.Output(user), nil
}
