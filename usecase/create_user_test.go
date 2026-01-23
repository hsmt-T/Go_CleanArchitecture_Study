package usecase_test

import (
	"go_cleanArchitecture_study/adapter/presenter"
	"go_cleanArchitecture_study/domain"
	"go_cleanArchitecture_study/usecase"
	"testing"
)


//RepositoryはMockを使う
type mockUserRepository struct {
	savedUser domain.User
}

func (m *mockUserRepository) Create(user domain.User) (domain.User, error) {
	m.savedUser = user
	return user, nil
}

func (m *mockUserRepository) FindByID(id domain.UserID) (domain.User, error) {
	return domain.User{}, nil
}

//Testコード

func TestCreteUser_Success(t *testing.T) {
	repo := &mockUserRepository{}
	presenter := presenter.NewCreateUserPresenter()

	uc := usecase.NewCreateUserInteractor(repo, presenter)

	input := usecase.CreateUserInput{
		Name: "test",
		Email: "test.gmail.com",
	}

	output, err := uc.Execute(input)

	if err != nil {
		t.Fatalf("入力えらー: %v", err)
	}

	if output.Name != "test" {
		t.Errorf("test 以外の Nameが入っている ：%v", output.Name)
	}

	if output.Email != "test.gmail.com" {
		t.Errorf("test.gmail.com 以外の Emailが入っている :%v", output.Email)
	}

	if output.ID == "" {
		t.Error("IDがない")
	}

	if output.CreatedAt.IsZero() {
		t.Error("CreatedAtがない")
	}
}

//わざとエラーが出るInputにして出るかTest
func TestCreateUser_ValidationError(t *testing.T) {
	repo := &mockUserRepository{}
	presenter := presenter.NewCreateUserPresenter()
	uc := usecase.NewCreateUserInteractor(repo, presenter)

	input := usecase.CreateUserInput{
		Name: "",
		Email: "",
	}

	_, err := uc.Execute(input)

	if err == nil {
		t.Fatal("バリデーションエラーがあるはずなのに出ていない")
	}
}