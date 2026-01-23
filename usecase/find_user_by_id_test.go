package usecase_test

import (
	"go_cleanArchitecture_study/adapter/presenter"
	"go_cleanArchitecture_study/domain"
	"go_cleanArchitecture_study/usecase"
	"testing"
	"time"
)

type mockFindByIDUserRepository struct {
	user domain.User
}

func (m *mockFindByIDUserRepository) Create(user domain.User) (domain.User, error) {
	return user, nil
}

func (m *mockFindByIDUserRepository) FindByID(id domain.UserID) (domain.User, error) {
	return m.user, nil
}

func TestFindByUser_Success (t *testing.T) {

	expectedUser := domain.NewUser(
		domain.UserID("123"),
		"test",
		"test@gmail.com",
		time.Now(),
	)

	repo := &mockFindByIDUserRepository{
		user: expectedUser,
	}
	presenter := presenter.NewFindUserPresenter()

	uc := usecase.NewFindUserByIDInteractor(repo, presenter)

	output, err := uc.Execute(usecase.FindUserByIDInput{
		ID: "123",
	})

	if err != nil {
		t.Fatalf("error %v", err)
	}

	if output.ID != "123" {
		t.Errorf("IDがmock通りじゃない %v", output.ID)
	}

	if output.Name != "test" {
		t.Errorf("Nameがmock通りじゃない %v", output.ID)
	}

	if output.Email != "test@gmail.com" {
		t.Errorf("Emailがmock通りじゃない %v", output.ID)
	}
}