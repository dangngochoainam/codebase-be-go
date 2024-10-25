package controller

import (
	"example/internal/common/helper/copyhepler"
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		CreateUser(ctx *gin.Context)
		CreateManyUser(ctx *gin.Context)
		FindOneUser(ctx *gin.Context)
		FindUsers(ctx *gin.Context)
		UpdateUserById(ctx *gin.Context)
		UpdateUser(ctx *gin.Context)
		SoftDeleteUser(ctx *gin.Context)
	}

	userController struct {
		userUseCase    usecase.UserUseCase
		modelConverter copyhepler.ModelConverter
	}
)

func NewUserController(userUseCase usecase.UserUseCase, modelConverter copyhepler.ModelConverter) UserController {
	return &userController{
		userUseCase:    userUseCase,
		modelConverter: modelConverter,
	}
}

func (u *userController) CreateUser(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	input := &dto.CreateUserRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.CreateUser(input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}

func (u *userController) CreateManyUser(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	input := &dto.CreateUsersRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.CreateUsers(input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}

func (u *userController) FindOneUser(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	input := &dto.FindOneUserRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.FindOneUser(input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}

func (u *userController) FindUsers(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	input := &dto.FindUsersRequestDTO{}
	err := appC.C.ShouldBindQuery(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.FindUsers(input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}

func (u *userController) UpdateUserById(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	id := appC.C.Param("id")
	input := &dto.UpdateUserRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.UpdateUserById(id, input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}

func (u *userController) UpdateUser(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	input := &dto.UpdateUserRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.UpdateUser(input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}

func (u *userController) SoftDeleteUser(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	id := appC.C.Param("id")

	// call usecase handle business
	data, err := u.userUseCase.SoftDeleteUser(id)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}
