package router

import (
	echoapi "go_cleanArchitecture_study/adapter/api/action/echo"
	"go_cleanArchitecture_study/adapter/presenter"
	"go_cleanArchitecture_study/adapter/repository"
	"go_cleanArchitecture_study/usecase"

	"github.com/labstack/echo/v4"
)

func StartEcho() {
	e := echo.New()

	repo := repository.NewUserMemoryRepository()
	presenter := presenter.NewCreateUserPresenter()
	uc := usecase.NewCreateUserInteractor(repo, presenter)
	handler := echoapi.NewCreateUserHandler(uc)
	
	e.POST("/users", handler.Handle)

	e.Start(":8080")
}