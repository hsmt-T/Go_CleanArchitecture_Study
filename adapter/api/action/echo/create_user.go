package echoapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go_cleanArchitecture_study/usecase"
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
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h CreateUserHandler) Handle(c echo.Context) error {
	var req createUserRequest

	// JSON → struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	// UseCase Input
	input := usecase.CreateUserInput{
		Name:  req.Name,
		Email: req.Email,
	}

	// UseCase 実行
	output, err := h.uc.Execute(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// レスポンス
	return c.JSON(http.StatusOK, output)
}
