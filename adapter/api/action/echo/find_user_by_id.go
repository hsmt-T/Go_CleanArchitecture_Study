package echoapi

import (
	"go_cleanArchitecture_study/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FindUserHandler struct {
	uc usecase.FindUserByIDUseCase
}

func NewFindUserHandler(uc usecase.FindUserByIDUseCase) FindUserHandler {
	return FindUserHandler{uc: uc}
}

func (h FindUserHandler) Handle(c echo.Context) error {
	id := c.Param("id")

	input := usecase.FindUserByIDInput{
		ID: id,
	}

	output, err := h.uc.Execute(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, output)
}
