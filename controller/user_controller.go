package controller

import (
	"go-api/cmd/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *userController {
	return &userController{
		UserUsecase: usecase,
	}
}

func (p *userController) GetUsers(ctx *gin.Context) {
	users, err := p.UserUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}

	ctx.JSON(http.StatusOK, users)
}
