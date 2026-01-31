package echoapi_test

import (
	"bytes"
	"encoding/json"
	"errors"
	echoapi "go_cleanArchitecture_study/adapter/api/action/echo"
	"go_cleanArchitecture_study/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// mock UseCase Success
type mockCreateUserUseCase struct {}

func (m *mockCreateUserUseCase) Execute(input usecase.CreateUserInput) (usecase.CreateUserOutput, error) {
	return usecase.CreateUserOutput{
		ID: "123",
		Name: 	input.Name,
		Email:  input.Email,
	}, nil
}

func TestCreateUserHandler_success(t *testing.T) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	uc := &mockCreateUserUseCase{}
	handler := echoapi.NewCreateUserHandler(uc)

	e.POST("/users", handler.Handle)

	reqBody := map[string]string{
		"name": "test",
		"email": "test@gmail.com",
	}
	body, _ := json.Marshal(reqBody)

	//ここで疑似的にPOST/usersをしている
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	t.Logf("response body: %s", rec.Body.String())

	if rec.Code != http.StatusOK {
		t.Fatalf("status code が違う: %v", rec.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("レスポンスjsonが壊れている: %v", err)
	}

	if res["name"] != "test" {
		t.Fatalf("nameが違う: %v", res["name"])
	}

	if res["email"] != "test@gmail.com" {
		t.Fatalf("emailが違う: %v", res["email"])
	}

	if res["created_at"] == "" {
		t.Fatalf("created_atがはいっていない: %v", res["created_at"])
	}
}

// mock UseCase Error
type mockCreateUserErrorUseCase struct {}

func (m *mockCreateUserErrorUseCase) Execute(input usecase.CreateUserInput) (usecase.CreateUserOutput, error) {
	return usecase.CreateUserOutput{}, errors.New("バリデーションエラー")
}

func TestCreateUserHandler_Error(t *testing.T) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	uc := &mockCreateUserErrorUseCase{}
	handler := echoapi.NewCreateUserHandler(uc)

	e.POST("/users", handler.Handle)

	//error用だからBody空
	reqBody := map[string]string{}

	body, _ := json.Marshal(reqBody)

	//ここで疑似的にPOST/usersをしている
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//実行
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("error なのに Statusが違う: %v", rec.Code)
	}
}

