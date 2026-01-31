package echoapi_test

import (
	"encoding/json"
	"errors"
	echoapi "go_cleanArchitecture_study/adapter/api/action/echo"
	"go_cleanArchitecture_study/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type mockFindByIDUseCase struct {}

func (m *mockFindByIDUseCase) Execute(input usecase.FindUserByIDInput) (usecase.FindUserByIDOutput, error ) {
	return usecase.FindUserByIDOutput{
		ID: input.ID,
		Name: "test",
		Email: "test@gmail.com",
		CreatedAt: time.Now(),
	}, nil
}

func TestFindByIDHandler_success(t *testing.T) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	uc := &mockFindByIDUseCase{}
	handler := echoapi.NewFindUserHandler(uc)

	e.GET("/users/:id", handler.Handle)

	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	t.Logf("body:%v", rec.Body.String())

	if rec.Code != http.StatusOK {
		t.Fatalf("Status Code が違う%v", rec.Code)
	}

	var res map[string]interface{}

	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("JSONが壊れている%v", err)
	}

	if res["id"] != "123" {
		t.Fatalf("idが違う %v", res["id"])
	}

	if res["name"] != "test" {
		t.Fatalf("nameが違う: %v", res["name"])
	}

	if res["email"] != "test@gmail.com" {
		t.Fatalf("emailが違う: %v", res["email"])
	}

	if res["createdAt"] == "" {
		t.Fatalf("created_atがはいっていない: %v", res["createdAt"])
	}
}


 //errorテスト用に新しくmockUsecaseを作る
 //このテストではmockUcが呼ばれているか確認するためにcalledのboolで確認
type mockFindByIDErrorUseCase struct {
	called bool
}

func (m *mockFindByIDErrorUseCase) Execute(input usecase.FindUserByIDInput) (usecase.FindUserByIDOutput, error) {
	m.called = true
	return usecase.FindUserByIDOutput{}, errors.New("error")
}

func TestFindByIDHandler_Error(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req,rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("")

	uc := &mockFindByIDErrorUseCase{}
	handler := echoapi.NewFindUserHandler(uc)

	err := handler.Handle(c)


	//assertはエラーがおきたときに終了せず最後までかくにんできる
	assert.NoError(t, err)

	//ucをよんだらだめな状況（Error）で読んでないかチェック
	assert.False(t, uc.called)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	body := rec.Body.String()
	assert.Contains(t,body, "error")
}

