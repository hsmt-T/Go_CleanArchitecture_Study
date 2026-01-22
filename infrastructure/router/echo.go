package router

import (
	echoapi "go_cleanArchitecture_study/adapter/api/action/echo"
	"go_cleanArchitecture_study/adapter/presenter"
	"go_cleanArchitecture_study/adapter/repository"
	"go_cleanArchitecture_study/infrastructure/database"
	"go_cleanArchitecture_study/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

func StartEcho() {
	e := echo.New()

	// ローカルDB
	// repo := repository.NewUserMemoryRepository()


	db, err := supabase.NewSupabase()

	if err != nil {
		log.Fatal("DB接続失敗", err)
	}

	repo := repository.NewUserSupabase(db)
	presenter := presenter.NewCreateUserPresenter()
	uc := usecase.NewCreateUserInteractor(repo, presenter)
	handler := echoapi.NewCreateUserHandler(uc)
	
	e.POST("/users", handler.Handle)

	e.Start(":8080")
}