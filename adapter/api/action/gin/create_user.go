package ginapi

import (
	"go_cleanArchitecture_study/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	uc usecase.CreateUserUseCase
}

func NewCreateUserHandler(uc usecase.CreateUserUseCase) CreateUserHandler {
	return CreateUserHandler{
		uc: uc,
	}
}

type createUserRequest struct {
	Name	string	`json:"name"`
	Email	string	`json:"email"`
}

func (h CreateUserHandler) Handle(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "無効なリクエスト",
		})
		return
	}

	input := usecase.CreateUserInput{
		Name: req.Name,
		Email: req.Email,
	}

	output, err := h.uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

//Gin は UseCase の interface しか知らない = クリーンアーキテクチャ