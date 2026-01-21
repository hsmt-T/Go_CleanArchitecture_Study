package router

import (
	ginapi "go_cleanArchitecture_study/adapter/api/action/gin"
	"go_cleanArchitecture_study/adapter/presenter"
	"go_cleanArchitecture_study/adapter/repository"
	"go_cleanArchitecture_study/usecase"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	r := gin.Default()

	repo := repository.NewUserMemoryRepository()
	presemter := presenter.NewCreateUserPresenter()
	uc := usecase.NewCreateUserInteractor(repo, presemter)
	handler := ginapi.NewCreateUserHandler(uc)

	r.POST("/users", handler.Handle)

	r.Run(":8080")
}