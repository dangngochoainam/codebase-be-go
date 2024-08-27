package controller

import (
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		GetUserList(ctx *gin.Context)
	}

	userController struct {
		userUseCase usecase.UserUseCase
	}
)

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: userUseCase,
	}
}

func (u *userController) GetUserList(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}
	// handle request input from client
	username := appC.C.Query("username")

	// call usecase handle business
	data, err := u.userUseCase.FindUsers(&dto.FindUsersRequestDTO{
		Username: username,
	})

	// handle response output into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}
