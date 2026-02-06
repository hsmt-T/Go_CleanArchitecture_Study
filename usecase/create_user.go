package usecase

import (
	"errors"
	"go_cleanArchitecture_study/domain"
	"time"
)

//型定義（今はここだけど後で細分化するのかな？）

type CreateUserInput struct {
	Name	string
	Email	string
}

type CreateUserOutput struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	CreatedAt	time.Time	`json:"created_at"`
}

// UseCase インターフェイス
// adapter（API)が知ることになる
type CreateUserUseCase interface {
	Execute(input CreateUserInput) (CreateUserOutput, error)
}

type CreateUserPresenter interface {
	Output(domain.User) CreateUserOutput
}

// domain に依存する/repository を使う が　抽象に依存
type createUserInteractor struct {
	userRepo domain.UserRepository 
	presenter CreateUserPresenter
	clock Clock
}

//コンストラクタ

func NewCreateUserInteractor(repo domain.UserRepository, presenter CreateUserPresenter, clock Clock) CreateUserUseCase {
	return &createUserInteractor{
		userRepo: repo,
		presenter: presenter,
		clock: clock,
	}
}

//メインの処理

func (i *createUserInteractor) Execute(input CreateUserInput) (CreateUserOutput,error) {

	//ここでバリデーション
	if input.Name == "" || input.Email == "" {
		return CreateUserOutput{}, errors.New("なまえとメールが未記入です")
	}

	now := i.clock.Now()

	//domainのUsers生成
	user := domain.NewUser(
		domain.NewUserID(),
		input.Name,
		input.Email,
		now,
	)

	// Repository（インターフェイス）に保存する
	createdUser, err := i.userRepo.Create(user)
	if err != nil {
		return CreateUserOutput{}, err
	}

	return i.presenter.Output(createdUser), nil
}
