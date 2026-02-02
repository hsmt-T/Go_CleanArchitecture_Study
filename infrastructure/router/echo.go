package router

import (
	"context"
	echoapi "go_cleanArchitecture_study/adapter/api/action/echo"
	"go_cleanArchitecture_study/adapter/presenter"
	"go_cleanArchitecture_study/adapter/repository"
	"go_cleanArchitecture_study/infrastructure/clock"
	"go_cleanArchitecture_study/infrastructure/database"
	"go_cleanArchitecture_study/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartEcho() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	

	// ローカルDB
	// repo := repository.NewUserMemoryRepository()

	//NeonDB
	db, err := database.NewNeon(context.Background())

	//supabase
	// db, err := database.NewSupabase()


	if err != nil {
		log.Fatal("DB接続失敗", err)
	}
	// supabaseRepo
	// repo := repository.NewUserSupabase(db)

	// neonRepo
	repo := repository.NewUserNeon(db)

	//create
	createPresenter := presenter.NewCreateUserPresenter()
	RealClock := clock.RealClock{}
	createUc := usecase.NewCreateUserInteractor(repo, createPresenter, RealClock)
	createHandler := echoapi.NewCreateUserHandler(createUc)

	//findByID
	findByIDPresenter := presenter.NewFindUserPresenter()
	findByIDUc := usecase.NewFindUserByIDInteractor(repo,findByIDPresenter)
	findByIDHandler := echoapi.NewFindUserHandler(findByIDUc)
	
	e.POST("/users", createHandler.Handle)
	e.GET("/users/:id", findByIDHandler.Handle)

	e.Start(":8080")
}